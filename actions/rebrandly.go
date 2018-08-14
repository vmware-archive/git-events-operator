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

func processGitHubMergeToMaster(e event.Event) error {
	merge := e.(*github.EventImplementation)
	fmt.Printf("Processing merge from %s\n", merge.CommiterEmail)
	return nil
}
