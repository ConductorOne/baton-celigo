package celigo

import (
	"fmt"
	"net/http"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
)

type Region int

const (
	BaseEUUrl = "https://api.eu.integrator.io"
	BaseUrl   = "https://api.integrator.io"

	USRegion Region = iota
	EURegion
)

type Client struct {
	uhttp.BaseHttpClient

	baseUrl string
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
		BaseHttpClient: *uhttp.NewBaseHttpClient(httpClient),
		baseUrl:        u,
	}, nil
}
