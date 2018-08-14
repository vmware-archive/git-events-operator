package action

import "github.com/heptiolabs/git-events-operator/event"

type ActionFunc func(event event.Event) error
