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

package cmd

import (
	"github.com/heptiolabs/git-events-operator/actions"
	"github.com/heptiolabs/git-events-operator/event"
	"github.com/heptiolabs/git-events-operator/event/github"
	"github.com/heptiolabs/git-events-operator/operator"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

// config.go
//
// This file is effectively the config file for the operator.
// In an order to keep version 1 simple, we make configuration changes
// here and then rebuild the image.
//
// We can pull these into configuration later, but for now this gives
// us a way to get started with relatively little overhead.

func init() {

	// Logger level
	// 4 Most verbose
	// 3 Preferred
	// 2 and 1 no recommended
	logger.Level = 3

}

var operatorConfig = &operator.Config{

	// Define which event brokers we want to use
	Brokers: []event.Broker{
		&github.EventBrokerImplementation{
			Owner: "heptio",
			Repo:  "advocacy",
			Path:  "content/event/",
		},
	},

	// Map specific event types to specific actions
	ActionMap: map[event.EventKind]action.ActionFunc{

		// Here we map NewFile to GenerateAndSendRebrandly link
		event.NewFile: action.GenerateAndSendRebrandlyLink,
	},
}
