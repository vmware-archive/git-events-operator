package rebrandly_go_sdk

import (
	"fmt"
	"os"
	"testing"

	"github.com/kubicorn/kubicorn/pkg/namer"
)

func TestCreateLink(t *testing.T) {
	client, err := NewRebrandlyClient()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := map[string]interface{}{
		// 8f104cc5b6ee4a4ba7897b06ac2ddcfb is the default rebrand.ly id
		"destination": fmt.Sprintf("http://%s.com", namer.RandomName()),
		"slashtag":    namer.RandomName(),
		"title":       "Test title",
		//"description": "A wonderful link",
		"domain": map[string]interface{}{
			"ref": "",
			"id":  "8f104cc5b6ee4a4ba7897b06ac2ddcfb",
		},
	}
	res, err := client.CreateLink(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if res.Response.Status != "200 OK" {
		t.Errorf("Unable to GET account")
	}
}

func TestListLinks(t *testing.T) {
	client, err := NewRebrandlyClient()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := map[string]interface{}{}
	res, err := client.ListLinks(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if res.Response.Status != "200 OK" {
		t.Errorf("Unable to GET account")
	}
}
