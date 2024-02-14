package connector

import (
	"context"
	"fmt"
	"io"

	"github.com/conductorone/baton-celigo/pkg/celigo"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/connectorbuilder"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
)

type Celigo struct {
	Client *celigo.Client
}

// ResourceSyncers returns a ResourceSyncer for each resource type that should be synced from the upstream service.
func (d *Celigo) ResourceSyncers(ctx context.Context) []connectorbuilder.ResourceSyncer {
	return []connectorbuilder.ResourceSyncer{
		newUserBuilder(d.Client),
	}
}

// Asset takes an input AssetRef and attempts to fetch it using the connector's authenticated http client
// It streams a response, always starting with a metadata object, following by chunked payloads for the asset.
func (d *Celigo) Asset(ctx context.Context, asset *v2.AssetRef) (string, io.ReadCloser, error) {
	return "", nil, nil
}

// Metadata returns metadata about the connector.
func (d *Celigo) Metadata(ctx context.Context) (*v2.ConnectorMetadata, error) {
	return &v2.ConnectorMetadata{
		DisplayName: "Celigo Baton Connector",
		Description: "The Celigo Baton Connector for the Celigo platform.",
	}, nil
}

// Validate is called to ensure that the connector is properly configured. It should exercise any API credentials
// to be sure that they are valid.
func (d *Celigo) Validate(ctx context.Context) (annotations.Annotations, error) {
	return nil, nil
}

// New returns a new instance of the connector.
func New(ctx context.Context, accessToken, region string) (*Celigo, error) {
	httpClient, err := uhttp.NewClient(
		ctx,
		uhttp.WithLogger(true, nil),
		uhttp.WithUserAgent("baton-celigo"),
	)
	if err != nil {
		return nil, err
	}

	var r celigo.Region
	switch region {
	case "us":
		r = celigo.USRegion
	case "eu":
		r = celigo.EURegion
	default:
		return nil, fmt.Errorf("invalid region: %s", region)
	}

	client := celigo.New(accessToken, r, httpClient)

	return &Celigo{
		Client: client,
	}, nil
}
