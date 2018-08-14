package cmd

import (
	"github.com/heptiolabs/git-events-operator/actions"
	"github.com/heptiolabs/git-events-operator/event"
	"github.com/heptiolabs/git-events-operator/event/github"
	"github.com/heptiolabs/git-events-operator/operator"
)

// config.go
//
// This file is effectively the config file for the operator.
// In an order to keep version 1 simple, we make configuration changes
// here and then rebuild the image.
//
// We can pull these into configuration later, but for now this gives
// us a way to get started with relatively little overhead.

var operatorConfig = &operator.Config{

	// Define which event brokers we want to use
	Brokers: []event.Broker{
		&github.EventBrokerImplementation{},
	},

	// Map specific event types to specific actions
	ActionMap: map[event.EventKind]action.ActionFunc{

		// Here we map MergeToMaster to GenerateAndSendRebrandly link
		event.MergeToMaster: action.GenerateAndSendRebrandlyLink,
	},
}
