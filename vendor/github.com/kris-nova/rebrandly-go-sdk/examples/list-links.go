package main

import (
	"fmt"
	"os"

	"github.com/kris-nova/rebrandly-go-sdk"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

func main() {

	// Set the logger level to debug the SDK if developing
	logger.Level = 4

	// Create a new client
	client, err := rebrandly_go_sdk.NewRebrandlyClient()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Use the client to send a GET request to the account endpoint
	params := map[string]interface{}{
		// 8f104cc5b6ee4a4ba7897b06ac2ddcfb is the default rebrand.ly id
		//"orderBy": "",
		//"slashtag":    "",
		//"title":       "Test title",
		//"description": "A wonderful link",
		"domain": map[string]interface{}{
			"ref": "",
			"id":  "8f104cc5b6ee4a4ba7897b06ac2ddcfb",
		},
	}
	res, err := client.ListLinks(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Format the response
	responseBody, err := res.Pretty()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Output
	fmt.Println(res.Response.Status)
	fmt.Println(responseBody)

}
