// Copyright © 2017 Heptio
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

	"time"

	"os"

	"strings"

	sp "github.com/SparkPost/gosparkpost"
	"github.com/heptiolabs/git-events-operator/event"
	"github.com/heptiolabs/git-events-operator/event/github"
	rebrandly "github.com/kris-nova/rebrandly-go-sdk"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

const (
	Domain                      = "rebrand.ly"
	SleepBeforeRepublishSeconds = 120
	HeptioAdvocacyEmail         = "advocacy@mailing.tgik8s.com"
	HeptioAdvocacyName          = "Heptio Advocacy"
)

// GenerateAndSendRebrandlyLink expects
func GenerateAndSendRebrandlyLink(e event.Event, q *event.Queue) error {

	// Map our support for the various kinds
	switch e.Kind() {
	case event.NewFile:
		return processNewFile(e, q)
	default:
		return fmt.Errorf("kind %s not supprted", e.Kind())
	}
	return nil
}

func processNewFile(e event.Event, q *event.Queue) error {
	switch e.Type() {
	case github.GithubImplementation:
		return processGitHubNewFile(e, q)
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

func processGitHubNewFile(e event.Event, q *event.Queue) error {

	// TODO @kris-nova this is the hackiest API parsing known to womankind. Please fix this, soon.
	newFile := e.(*github.EventImplementation)

	// Calculate the expected shortUrl
	slashtag := strings.Replace(rebrandlyHash(newFile.FileName), ".md", "", 1)
	expectedShortURL := fmt.Sprintf("%s/%s", Domain, slashtag)

	// -----------------------------------------------------------------------------------------------------------------
	//
	// ENSURE REBRANDLY LINK HERE
	//
	// -----------------------------------------------------------------------------------------------------------------
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

		params := map[string]interface{}{
			"destination": fmt.Sprintf("https://advocacy.heptio.com/event/%s", newFile.FileName),
			"slashtag":    slashtag,
			"title":       fmt.Sprintf("[Heptio Advocacy] Automatic link created for page [%s]", newFile.FileName),
			//"description": "A wonderful link",
			"domain": map[string]interface{}{
				"ref": "",
				"id":  "8f104cc5b6ee4a4ba7897b06ac2ddcfb",
			},
		}
		res, err := client.CreateLink(params)
		if err != nil {
			return fmt.Errorf("unable to create new link: %v", err)
		}
		if res.Response.StatusCode == 200 {
			logger.Info("Status Code [%d %s]", res.Response.StatusCode, res.Response.Status)
		} else {
			// Republish event and try again

			logger.Info("Failure creating link, republishing event in queue")
			logger.Info("Response code [%s]", res.Response.Status)
			body, err := res.Body()
			if err == nil {
				logger.Info("Output dump: %s", body)
			}
			time.Sleep(time.Duration(time.Second * SleepBeforeRepublishSeconds))
			q.AddEvent(e)
			return nil
		}
		for authorEmail, authorName := range newFile.Authors {
			// ---------------------------------------------------------------------------------------------------------
			//
			// SEND EMAIL HERE
			//
			// ---------------------------------------------------------------------------------------------------------
			logger.Info("Alerting user [%s] via email [%s] of new link [%s] for file [%s]", authorName, authorEmail, expectedShortURL, newFile.Name)
			apiKey := os.Getenv("SPARKPOST_API_KEY")
			cfg := &sp.Config{
				BaseUrl:    "https://api.sparkpost.com",
				ApiKey:     apiKey,
				ApiVersion: 1,
			}
			var client sp.Client
			err := client.Init(cfg)
			if err != nil {
				return fmt.Errorf("unable to create sparkpost client: %v", err)
			}

			tx := &sp.Transmission{
				Recipients: []string{authorEmail},
				Content: sp.Content{
					HTML: getHTMLEmail(authorName, newFile.FileName, expectedShortURL),
					From: map[string]string{
						"name":  HeptioAdvocacyName,
						"email": HeptioAdvocacyEmail,
					},
					Subject: fmt.Sprintf("[Heptio Advocacy] Your new event [%s]", newFile.FileName),
				},
			}
			_, _, err = client.Send(tx)
			if err != nil {
				return fmt.Errorf("unable to send email: %v", err)
			}

		}

	} else {
		logger.Info("Link exists [%s] - bypassing", newFile.FileName)
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
