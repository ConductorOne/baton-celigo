//go:build ignore

package main

import (
	"github.com/conductorone/baton-celigo/pkg/config"
	cfg "github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	cfg.Generate("celigo", config.ConfigurationFields, config.FieldRelationships)
}
