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
	res, err := client.GetAccount()
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
