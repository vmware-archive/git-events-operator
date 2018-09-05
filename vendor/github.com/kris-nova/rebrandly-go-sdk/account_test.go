package rebrandly_go_sdk

import (
	"fmt"
	"os"
	"testing"
)

func TestGetAccount(t *testing.T) {
	client, err := NewRebrandlyClient()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	res, err := client.GetAccount()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if res.Response.Status != "200 OK" {
		t.Errorf("Unable to GET account")
	}

}
