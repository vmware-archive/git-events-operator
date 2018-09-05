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

	"encoding/json"

	"hash/fnv"

	"strconv"

	"github.com/heptiolabs/git-events-operator/event"
	"github.com/heptiolabs/git-events-operator/event/github"
	rebrandly "github.com/kris-nova/rebrandly-go-sdk"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

const (
	Domain = "rebrand.ly"
)

// GenerateAndSendRebrandlyLink expects
func GenerateAndSendRebrandlyLink(e event.Event) error {

	// Map our support for the various kinds
	switch e.Kind() {
	case event.NewFile:
		return processNewFile(e)
	default:
		return fmt.Errorf("kind %s not supprted", e.Kind())
	}
	return nil
}

func processNewFile(e event.Event) error {
	switch e.Type() {
	case github.GithubImplementation:
		return processGitHubNewFile(e)
	default:
		return fmt.Errorf("unknown implementation: %s", e.Type())
	}
	return nil
}

//
//type RebrandlyLink struct {
//	id          string
//	linkId      int
//	title       string
//	slashtag    string
//	destination string
//	createdAt   string
//	updatedAt   string
//	status      string
//	clicks      int
//	isPublic    bool
//	shortUrl    string
//	domainId    string
//	domainName  string
//	domain      struct {
//		id       string
//		ref      string
//		fullName string
//		active   bool
//	}
//	creator struct {
//		id        string
//		fullName  string
//		avatarUrl string
//	}
//	integrated bool
//}

func processGitHubNewFile(e event.Event) error {

	// TODO @kris-nova this is the hackiest API parsing known to womankind. Please fix this, soon.
	newFile := e.(*github.EventImplementation)

	// Calculate the expected shortUrl
	expectedShortURL := fmt.Sprintf("%s/%s", Domain, rebrandlyHash(newFile.FileName))

	// Connect with Rebrandly
	client, err := rebrandly.NewRebrandlyClient()
	if err != nil {
		return fmt.Errorf("unable to auth with Rebrandly: %v", err)
	}
	response, err := client.ListLinks(make(map[string]interface{}))
	body, err := response.Body()
	if err != nil {
		return fmt.Errorf("unable to parse response from Rebrandly: %v", err)
	}
	var links []interface{}
	bodyBytes := []byte(body)
	json.Unmarshal(bodyBytes, &links)
	found := false
	for _, link := range links {
		linkMap := link.(map[string]interface{})
		//slashtag := linkMap["slashtag"]
		shortUrl := linkMap["shortUrl"]
		if shortUrl == expectedShortURL {
			found = true
			break
		}
	}
	if !found {
		logger.Info("Creating new rebrandly link [%s] for file [%s]", expectedShortURL, newFile.FileName)
		// TODO Create link
		for authorEmail, authorName := range newFile.Authors {
			logger.Info("Alerting user [%s] via email [%s] of new link [%s] for file [%s]", authorName, authorEmail, expectedShortURL, newFile.Name)
			// TODO Email user
		}

	}

	return nil
}

func rebrandlyHash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	ui32 := h.Sum32()
	i := int(ui32)
	str := strconv.Itoa(i)
	return str
}
