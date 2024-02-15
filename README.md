![Baton Logo](./docs/images/baton-logo.png)

# `baton-celigo` [![Go Reference](https://pkg.go.dev/badge/github.com/conductorone/baton-celigo.svg)](https://pkg.go.dev/github.com/conductorone/baton-celigo) ![main ci](https://github.com/conductorone/baton-celigo/actions/workflows/main.yaml/badge.svg)

`baton-celigo` is a connector for Baton built using the [Baton SDK](https://github.com/conductorone/baton-sdk). It works with Celigo API.

Check out [Baton](https://github.com/conductorone/baton) to learn more about the project in general.

# Prerequisites

Connector requires bearer access token that is used throughout the communication with API. To obtain this token, you have to create one in Celigo. More in information about how to generate token [here](https://docs.celigo.com/hc/en-us/articles/360042281231-Getting-started-with-standard-REST-API)). 

After you have obtained access token, you can use it with connector. You can do this by setting `BATON_CELIGO_ACCESS_TOKEN` or by passing `--celigo-access-token`.

Also you can set up region. Either US (default) or EU. It's can be done by `BATON_REGION` or by passing `--region`. 

# Getting Started

## brew

```
brew install conductorone/baton/baton conductorone/baton/baton-celigo

BATON_CELIGO_ACCESS_TOKEN=token BATON_REGION=eu baton-celigo
baton resources
```

## docker

```
docker run --rm -v $(pwd):/out -e BATON_CELIGO_ACCESS_TOKEN=token -e BATON_REGION=eu ghcr.io/conductorone/baton-celigo:latest -f "/out/sync.c1z"
docker run --rm -v $(pwd):/out ghcr.io/conductorone/baton:latest -f "/out/sync.c1z" resources
```

## source

```
go install github.com/conductorone/baton/cmd/baton@main
go install github.com/conductorone/baton-celigo/cmd/baton-celigo@main

BATON_CELIGO_ACCESS_TOKEN=token BATON_REGION=eu baton-celigo
baton resources
```

# Data Model

`baton-celigo` will fetch information about the following Baton resources:

- Users
- Integrations
- Roles (user's access level)

# Contributing, Support and Issues

We started Baton because we were tired of taking screenshots and manually building spreadsheets. We welcome contributions, and ideas, no matter how small -- our goal is to make identity and permissions sprawl less painful for everyone. If you have questions, problems, or ideas: Please open a Github Issue!

See [CONTRIBUTING.md](https://github.com/ConductorOne/baton/blob/main/CONTRIBUTING.md) for more details.

# `baton-celigo` Command Line Usage

```
baton-celigo

Usage:
  baton-celigo [flags]
  baton-celigo [command]

Available Commands:
  capabilities       Get connector capabilities
  completion         Generate the autocompletion script for the specified shell
  help               Help about any command

Flags:
      --celigo-access-token string   Celigo Access Token
      --client-id string             The client ID used to authenticate with ConductorOne ($BATON_CLIENT_ID)
      --client-secret string         The client secret used to authenticate with ConductorOne ($BATON_CLIENT_SECRET)
  -f, --file string                  The path to the c1z file to sync with ($BATON_FILE) (default "sync.c1z")
  -h, --help                         help for baton-celigo
      --log-format string            The output format for logs: json, console ($BATON_LOG_FORMAT) (default "json")
      --log-level string             The log level: debug, info, warn, error ($BATON_LOG_LEVEL) (default "info")
  -p, --provisioning                 This must be set in order for provisioning actions to be enabled. ($BATON_PROVISIONING)
      --region string                Region (default "us")
  -v, --version                      version for baton-celigo

Use "baton-celigo [command] --help" for more information about a command.
```