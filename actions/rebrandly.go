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

package action

import (
	"fmt"

	"github.com/heptiolabs/git-events-operator/event"
	"github.com/heptiolabs/git-events-operator/event/github"
)

// GenerateAndSendRebrandlyLink expects
func GenerateAndSendRebrandlyLink(e event.Event) error {

	// Map our support for the various kinds
	switch e.Kind() {
	case event.MergeToMaster:
		return processMergeToMaster(e)
	default:
		return fmt.Errorf("kind %s not supprted", e.Kind())
	}
	return nil
}

// processMergeToMaster will handle
func processMergeToMaster(e event.Event) error {
	switch e.Type() {
	case github.GithubImplementation:
		return processGitHubMergeToMaster(e)
	default:
		return fmt.Errorf("unknown implementation: %s", e.Type())
	}
	return nil
}

// processGitHubMergeToMaster will process a merge that is specific to GitHub
func processGitHubMergeToMaster(e event.Event) error {
	merge := e.(*github.EventImplementation)
	fmt.Printf("Processing merge from %s\n", merge.Name)
	return nil
}
