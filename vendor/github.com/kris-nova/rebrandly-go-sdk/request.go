package rebrandly_go_sdk

import (
	"fmt"
	"net/http"

	"bytes"
	"encoding/json"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

type rebrandlyHTTPMethod string

const (
	GET          rebrandlyHTTPMethod = "get"
	POST         rebrandlyHTTPMethod = "post"
	rebrandlyAPI                     = "https://api.rebrandly.com"
)

type rebrandlyRequest struct {
	method   rebrandlyHTTPMethod
	endpoint string
	apiKey   string
	params   rebrandlyParamters
}

type rebrandlyParamters map[string]interface{}

// execute is the core of the SDK request engine. Here we hard
// code a lot of logic that is specific to rebrandly.
func (res *rebrandlyRequest) execute() (*http.Response, error) {
	if res.method == "" {
		return nil, fmt.Errorf("empty method for request")
	}
	if res.endpoint == "" {
		return nil, fmt.Errorf("empty endpoint for request")
	}
	if res.apiKey == "" {
		return nil, fmt.Errorf("empty apiKey for request")
	}

	// URI is the formatted resolvable uri for the request.
	// Use this in your request.
	// TODO: @kris-nova can we use a proper Join() here instead of Sprintf()
	uri := fmt.Sprintf("%s%s", rebrandlyAPI, res.endpoint)
	logger.Info("path: %s", uri)
	switch res.method {

	// GET will send a GET request, here we hard code Content-Type and apiKey headers.
	case GET:

		// Create the body string, we hard code JSON
		jsonBodyBytes, err := json.Marshal(res.params)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON body: %v", err)
		}

		// Build the request
		client := &http.Client{}
		bodyBuffer := &bytes.Buffer{}
		bodyBuffer.Write(jsonBodyBytes)

		req, err := http.NewRequest("get", uri, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to create new GET request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("apiKey", res.apiKey)
		return client.Do(req)
	case POST:

		// Create the body string, we hard code JSON
		jsonBodyBytes, err := json.Marshal(res.params)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON body: %v", err)
		}

		// Build the request
		client := &http.Client{}
		bodyBuffer := &bytes.Buffer{}
		bodyBuffer.Write(jsonBodyBytes)
		req, err := http.NewRequest("post", uri, bodyBuffer)
		if err != nil {
			return nil, fmt.Errorf("unable to create new POST request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("apiKey", res.apiKey)
		return client.Do(req)
	default:
		return nil, fmt.Errorf("missing or invalid method for request")
	}

	return nil, nil

}
