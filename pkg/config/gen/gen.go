package main

import (
	cfg "github.com/conductorone/baton-celigo/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("celigo", cfg.Config)
}
