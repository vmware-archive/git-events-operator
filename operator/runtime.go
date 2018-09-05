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

package operator

import (
	"reflect"

	"runtime"

	"github.com/heptiolabs/git-events-operator/actions"
	"github.com/heptiolabs/git-events-operator/event"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

type Config struct {
	Brokers   []event.Broker
	ActionMap map[event.EventKind]action.ActionFunc
}

// Reconcile is a never ending loop that will attempt to reconcile
// events and action for the operator.
func Reconcile(cfg *Config) error {

	queue := event.NewQueue(cfg.Brokers)
	logger.Info("Loading event queue...")
	// Watch for errors from the queue concurrently
	go func() {
		errch := queue.ConcurrentStart()
		logger.Info("Starting even brokers...")
		for {
			err := <-errch
			// TODO do we want to break on error?
			logger.Warning("Error from message queue: %v", err)
		}
	}()

	for kind, f := range cfg.ActionMap {
		logger.Info("Mapping [%v]   :   [%s]", kind, runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	}

	for {
		event, err := queue.PopFromQueue()
		if err != nil {
			// Nothing in queue, don't log
			continue
		}

		// Logic to map EventKind's to ActionFunc's
		kind := event.Kind()
		actionMap := cfg.ActionMap
		action, ok := actionMap[kind]
		if !ok {

			if kind == "" {
				logger.Warning("Empty Kind for event, ignoring.")
				continue
			}

			// Kind not mapped
			logger.Warning("Kind %d not mapped in configuration, ignoring event", kind)
			continue
		}

		// Here we have an action and an event, woo!
		// Call the action
		err = action(event)
		if err != nil {
			logger.Warning("Unable to complete action: %v", err)
		}

	}
	return nil
}
