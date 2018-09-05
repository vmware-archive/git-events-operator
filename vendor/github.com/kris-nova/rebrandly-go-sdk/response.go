package rebrandly_go_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type rebrandlyResponse struct {
	Response *http.Response
}

func (r *rebrandlyResponse) Body() (string, error) {
	bodyData, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		return "", fmt.Errorf("unable to parse body; %v", err)
	}
	return string(bodyData), nil
}

func (r *rebrandlyResponse) Pretty() (string, error) {
	bodyData, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		return "", fmt.Errorf("unable to parse body; %v", err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, bodyData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("unable to json ident: %v", err)
	}
	formattedBytes := out.Bytes()
	return string(formattedBytes), nil
}
