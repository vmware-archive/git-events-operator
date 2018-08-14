package operator

import (
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
	// Watch for errors from the queue concurrently
	go func() {
		errch := queue.ConcurrentStart()
		for {
			err := <-errch
			// TODO do we want to break on error?
			logger.Warning("Error from message queue: %v", err)
		}
	}()
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
