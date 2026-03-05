package config

import (
	"context"
	"errors"

	"github.com/conductorone/baton-sdk/pkg/field"
)

var (
	CeligoAccessTokenField = field.StringField(
		"celigo-access-token",
		field.WithRequired(true),
		field.WithIsSecret(true),
		field.WithDescription("Celigo Access Token"),
	)
	RegionField = field.StringField(
		"region",
		field.WithDefaultValue("us"),
		field.WithDescription("Region"),
	)
	BaseURLField = field.StringField(
		"base-url",
		field.WithDescription("Override the Celigo API URL (for testing)"),
		field.WithHidden(true),
		field.WithExportTarget(field.ExportTargetCLIOnly),
	)

	// ConfigurationFields defines all configurable fields for the connector.
	ConfigurationFields = []field.SchemaField{
		CeligoAccessTokenField,
		RegionField,
		BaseURLField,
	}

	// FieldRelationships defines relationships between fields (constraints).
	FieldRelationships = []field.SchemaFieldRelationship{}
)

//go:generate go run ./gen
var Config = field.NewConfiguration(ConfigurationFields)

// ValidateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func ValidateConfig(ctx context.Context, cfg *Celigo) error {
	if cfg.CeligoAccessToken == "" {
		return errors.New("celigo-access-token is required")
	}
	return nil
}
