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
	"time"

	"github.com/heptiolabs/git-events-operator/event"
)

const (
	GithubImplementation event.ImplementationType = "GitHub"
)

type EventImplementation struct {
	Name               string
	KindValue          event.EventKind
	ImplementationType event.ImplementationType
	CommiterEmail      string
}

func (e *EventImplementation) Kind() event.EventKind {
	return e.KindValue
}

func (e *EventImplementation) Type() event.ImplementationType {
	return e.ImplementationType
}

type EventBrokerImplementation struct {
	queue *event.Queue
}

func (b *EventBrokerImplementation) ConcurrentWatch(queue *event.Queue) chan error {
	b.queue = queue
	errch := make(chan error)
	go func() {
		// Concurrent logic for GitHub goes here
		// return errors over the channel

		// Update queue with events as they come in
		// Here we simulate some
		//
		// TODO remove these and have real events
		//
		b.queue.AddEvent(&EventImplementation{
			Name:               "A fabulous event",
			KindValue:          event.MergeToMaster,
			ImplementationType: GithubImplementation,
			CommiterEmail:      "kris@nivenly.com",
		})

		time.Sleep(1 * time.Second)

		b.queue.AddEvent(&EventImplementation{
			Name:               "Another fabulous event",
			KindValue:          event.MergeToMaster,
			ImplementationType: GithubImplementation,
			CommiterEmail:      "kris@nivenly.com",
		})

		time.Sleep(1 * time.Second)

		b.queue.AddEvent(&EventImplementation{
			Name:               "The last fabulous event",
			KindValue:          event.MergeToMaster,
			ImplementationType: GithubImplementation,
			CommiterEmail:      "kris@nivenly.com",
		})

	}()
	return errch
}
