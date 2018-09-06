package rebrandly_go_sdk

import (
	"os"
	"fmt"
)

const (
	REBRANDLY_API_KEY_VARIABLE_NAME="REBRANDLY_API_KEY"
)

type rebrandlyClient struct {
	apiKey string
}


// NewRebrandlyClient authenticates a new Rebrandly client. This
// function expects the environmental variable `REBRANDLY_API_KEY`
// to be set, or will error.
func NewRebrandlyClient() (*rebrandlyClient, error) {

	apiKey := os.Getenv(REBRANDLY_API_KEY_VARIABLE_NAME)
	if apiKey == "" {
		return nil, fmt.Errorf("empty api key, please set REBRANDLY_API_KEY")
	}
	if len(apiKey) != 32 {
		return nil, fmt.Errorf("invalid api key, key must be 32 characters long")
	}
	return &rebrandlyClient{
		apiKey: apiKey,
	}, nil
}





