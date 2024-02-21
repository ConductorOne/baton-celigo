package celigo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Region int

const (
	BaseEUUrl = "https://api.eu.integrator.io"
	BaseUrl   = "https://api.integrator.io"

	USRegion Region = iota
	EURegion
)

type Client struct {
	httpClient  *http.Client
	accessToken string
	baseUrl     string
}

func New(accessToken string, region Region, httpClient *http.Client) (*Client, error) {
	var u string
	switch region {
	case USRegion:
		u = BaseUrl
	case EURegion:
		u = BaseEUUrl
	default:
		return nil, fmt.Errorf("invalid region: %d", region)
	}

	return &Client{
		httpClient:  httpClient,
		accessToken: accessToken,
		baseUrl:     u,
	}, nil
}

func (c *Client) newRequestWithDefaultHeaders(ctx context.Context, method string, url *url.URL, body ...interface{}) (*http.Request, error) {
	var buffer io.ReadWriter
	if body != nil {
		buffer = new(bytes.Buffer)
		err := json.NewEncoder(buffer).Encode(body[0])

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c *Client) do(req *http.Request, response ...interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response != nil {
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(response[0])
	}

	return resp, err
}
