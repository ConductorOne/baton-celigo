package celigo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	BaseEUUrl = "https://api.eu.integrator.io"
	BaseUrl   = "https://api.integrator.io"
)

type Client struct {
	httpClient  *http.Client
	accessToken string
	baseUrl     string
}

func New(accessToken string, EURegion bool, httpClient *http.Client) *Client {
	baseUrl := BaseUrl
	if EURegion {
		baseUrl = BaseEUUrl
	}

	return &Client{
		httpClient:  httpClient,
		accessToken: accessToken,
		baseUrl:     baseUrl,
	}
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

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return resp, err
	}

	if response != nil {
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(response[0])
	}

	return resp, err
}
