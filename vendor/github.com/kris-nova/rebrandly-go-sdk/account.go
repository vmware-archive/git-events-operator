package rebrandly_go_sdk

func (c *rebrandlyClient) GetAccount() (*rebrandlyResponse, error) {
	req := &rebrandlyRequest{
		method:   GET,
		endpoint: "/v1/account",
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
