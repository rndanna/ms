package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseAPI struct {
	Data []byte  `json:"data"`
	Err  *string `json:"err"`
}

type BaseClient struct {
	HTTPClient *http.Client
}

func (c *BaseClient) SendRequest(req *http.Request) (resp []byte, err error) {
	if c.HTTPClient == nil {
		return resp, fmt.Errorf("no http client")
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return resp, fmt.Errorf("failed to send request. error: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("status code !ok")
	}
	var responseAPI ResponseAPI
	if err = json.NewDecoder(response.Body).Decode(&responseAPI); err == nil {
		return resp, fmt.Errorf("failed  decode. error: %w", err)
	}

	return responseAPI.Data, err
}

func (c *BaseClient) Close() error {
	c.HTTPClient = nil
	return nil
}
