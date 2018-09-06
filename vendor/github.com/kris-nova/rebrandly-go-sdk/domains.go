package rebrandly_go_sdk

import "fmt"

func (c *rebrandlyClient) ListDomains() (*rebrandlyResponse, error) {
	req := &rebrandlyRequest{
		method:   GET,
		endpoint: "/v1/domains",
		apiKey:   c.apiKey,
	}
	resp, err := req.execute()
	if err != nil {
		return nil, err
	}
	return &rebrandlyResponse{
		Response: resp,
	}, nil
}

func (c *rebrandlyClient) GetDomain(params rebrandlyParamters) (*rebrandlyResponse, error) {
	req := &rebrandlyRequest{
		method:   GET,
		endpoint: fmt.Sprintf("/v1/domains"),
		apiKey:   c.apiKey,
		params:   params,
	}
	resp, err := req.execute()
	if err != nil {
		return nil, err
	}
	return &rebrandlyResponse{
		Response: resp,
	}, nil
}
