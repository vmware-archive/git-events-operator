package rebrandly_go_sdk

import "fmt"

func (c *rebrandlyClient) CreateLink(params rebrandlyParamters) (*rebrandlyResponse, error) {
	req := &rebrandlyRequest{
		method:   POST,
		endpoint: fmt.Sprintf("/v1/links/"),
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

func (c *rebrandlyClient) ListLinks(params rebrandlyParamters) (*rebrandlyResponse, error) {
	req := &rebrandlyRequest{
		method:   GET,
		endpoint: fmt.Sprintf("/v1/links/"),
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
