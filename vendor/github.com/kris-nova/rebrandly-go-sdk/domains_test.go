package rebrandly_go_sdk

import (
	"fmt"
	"os"
	"testing"
)

func TestListDomains(t *testing.T) {
	client, err := NewRebrandlyClient()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	res, err := client.ListDomains()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if res.Response.Status != "200 OK" {
		t.Errorf("Unable to GET account")
	}
}

func TestGetDomain(t *testing.T) {
	client, err := NewRebrandlyClient()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := map[string]interface{}{
		// 8f104cc5b6ee4a4ba7897b06ac2ddcfb is the default rebrand.ly id
		"id": "8f104cc5b6ee4a4ba7897b06ac2ddcfb",
	}
	res, err := client.GetDomain(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if res.Response.Status != "200 OK" {
		t.Errorf("Unable to GET account")
	}
}
