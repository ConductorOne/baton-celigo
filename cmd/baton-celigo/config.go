package main

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-sdk/pkg/cli"
	"github.com/spf13/cobra"
)

// config defines the external configuration required for the connector to run.
type config struct {
	cli.BaseConfig `mapstructure:",squash"` // Puts the base config options in the same place as the connector options

	CeligoAccessToken string `mapstructure:"celigo-access-token"`
	Region            string `mapstructure:"region"`
}

// validateConfig is run after the configuration is loaded, and should return an error if it isn't valid.
func validateConfig(ctx context.Context, cfg *config) error {
	if cfg.CeligoAccessToken == "" {
		return fmt.Errorf("celigo-access-token is required")
	}

	return nil
}

func cmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("celigo-access-token", "", "Celigo Access Token")
	cmd.PersistentFlags().String("region", "us", "Region")
}
