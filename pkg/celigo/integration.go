package celigo

import (
	"context"
	"net/http"
	"net/url"
)

type Integration struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Client) ListIntegrations(ctx context.Context, nextPageLink string) ([]Integration, string, *http.Response, error) {
	var err error
	stringUrl := nextPageLink
	if stringUrl == "" {
		stringUrl, err = url.JoinPath(c.baseUrl, "/v1/integrations")
		if err != nil {
			return nil, "", nil, err
		}
	}

	u, err := url.Parse(stringUrl)
	if err != nil {
		return nil, "", nil, err
	}

	req, err := c.newRequestWithDefaultHeaders(ctx, http.MethodGet, u)
	if err != nil {
		return nil, "", nil, err
	}

	var integrations []Integration
	resp, err := c.do(req, &integrations)
	if err != nil {
		return nil, "", nil, err
	}

	newNextPageLink := resp.Header.Get("Link")

	return integrations, newNextPageLink, resp, nil
}
