package celigo

import (
	"context"
	"net/http"
	"net/url"
)

type (
	IntegrationsAccessLevel struct {
		AccessLevel   string `json:"accessLevel"`
		IntegrationId string `json:"_integrationId"`
	}
	UserDetails struct {
		Email             string `json:"email"`
		AllowedToResetMFA bool   `json:"allowedToResetMFA"`
		AccountSSOLinked  string `json:"accountSSOLinked"`
	}
	User struct {
		Id                       string                    `json:"_id"`
		AccessLevel              string                    `json:"accessLevel"`
		IntegrationsAccessLevels []IntegrationsAccessLevel `json:"integrationAccessLevel"`
		AccountSSORequired       bool                      `json:"accountSSORequired"`
		Details                  UserDetails               `json:"sharedWithUser"`
	}
)

func (c *Client) ListUsers(ctx context.Context, nextPageLink string) ([]User, string, *http.Response, error) {
	var err error
	stringUrl := nextPageLink
	if stringUrl == "" {
		stringUrl, err = url.JoinPath(c.baseUrl, "/v1/ashares")
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

	var users []User
	resp, err := c.do(req, &users)
	if err != nil {
		return nil, "", nil, err
	}

	newNextPageLink := resp.Header.Get("Link")

	return users, newNextPageLink, resp, nil
}
