// Copyright © 2017 The Kubicorn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package github

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/heptiolabs/git-events-operator/event"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

const (
	GithubImplementation      event.ImplementationType = "GitHub"
	LoopIterationSleepSeconds                          = 3
)

type EventImplementation struct {
	Name               string
	KindValue          event.EventKind
	ImplementationType event.ImplementationType
	Authors            map[string]string
	FileName           string
}

func (e *EventImplementation) Kind() event.EventKind {
	return e.KindValue
}

func (e *EventImplementation) Type() event.ImplementationType {
	return e.ImplementationType
}

type EventBrokerImplementation struct {
	client *github.Client
	queue  *event.Queue

	// Path Information
	// TODO @kris-nova this might be the wrong place to define this if we want more than 1 instance but works for now

	Owner string
	Repo  string
	Path  string
}

// ConcurrentWatch will watch a directory and send an event for each file in the directory, and when a new file is created.
func (b *EventBrokerImplementation) ConcurrentWatch(queue *event.Queue) chan error {
	b.queue = queue
	errch := make(chan error)
	go func() {
		// Concurrent logic for GitHub goes here
		// return errors over the channel

		var cache []string
		for {
			time.Sleep(time.Duration(time.Second * LoopIterationSleepSeconds))
			logger.Debug("Querying repository")

			opt := &github.RepositoryContentGetOptions{}
			_, dirContent, _, err := b.client.Repositories.GetContents(context.TODO(), b.Owner, b.Repo, b.Path, opt)
			if err != nil {
				errch <- fmt.Errorf("unable to download contents from repository: %v", err)
				return
			}
			for _, repoContent := range dirContent {
				cached := false
				name := *repoContent.Name
				//fmt.Printf("(%s)\n", name)
				for _, cachedFile := range cache {
					if name == cachedFile {
						cached = true
						break
					}
				}
				if cached {
					continue
				}

				// Hooray a new file, let's look up the author information
				//opt := &github.ListContributorsOptions{
				//	ListOptions: github.ListOptions{},
				//}
				logger.Info("Parsing file: %s", name)
				opt := &github.CommitsListOptions{
					Path: fmt.Sprintf("%s%s", b.Path, name),
				}
				commits, _, err := b.client.Repositories.ListCommits(context.TODO(), b.Owner, b.Repo, opt)
				if err != nil {
					errch <- fmt.Errorf("unable to list commits: %v", err)
					return
				}
				//logger.Info("Number of commits: %d", len(commits))
				authors := make(map[string]string)
				for _, commit := range commits {

					// TODO This is very strange logic and if we are having issues with empty/missing emails
					// or email in general it's probably happening here.
					author := commit.Commit.Author
					if author == nil {
						continue
					}

					// Get the name and email of the committer
					name := author.GetName()
					email := author.GetEmail()

					//fmt.Printf("%+v\n", commit)
					//fmt.Println(name, email)

					// Index on email so we know we don't spam anyones inbox
					authors[email] = name
				}

				// Send the event to the queue
				event := &EventImplementation{
					Name:               name,
					KindValue:          event.NewFile,
					ImplementationType: GithubImplementation,
					Authors:            authors,
					FileName:           name,
				}
				queue.AddEvent(event)
				logger.Info("Adding event: %s", name)

				// Cache this file
				cache = append(cache, name)

			}
		}

	}()
	return errch
}

func (b *EventBrokerImplementation) Auth() error {
	user := os.Getenv("GITHUB_USER")
	if user == "" {
		return fmt.Errorf("missing or empty GITHUB_USER environmental variable, unable to authenticate with GitHub")
	}
	pass := os.Getenv("GITHUB_PASS")
	if user == "" {
		return fmt.Errorf("missing or empty GITHUB_PASS environmental variable, unable to authenticate with GitHub")
	}
	authTransport := github.BasicAuthTransport{
		Username: user,
		Password: pass,
	}
	client := github.NewClient(authTransport.Client())
	logger.Info("Authenticated with GitHub [%s]", user)
	b.client = client
	return nil
}
