# Caddy UI server

➡️ *Server responsible to serve the Caddy UI and communicate with the Caddy server API*

Please refer to the [**caddy-ui**](https://github.com/qdm12/caddy-ui) repository for end-user instructions.

[![Build status](https://github.com/qdm12/caddy-ui-server/workflows/Buildx%20latest/badge.svg)](https://github.com/qdm12/caddy-ui-server/actions?query=workflow%3A%22Buildx+latest%22)
[![Join Slack channel](https://img.shields.io/badge/slack-@qdm12-yellow.svg?logo=slack)](https://join.slack.com/t/qdm12/shared_invite/enQtOTE0NjcxNTM1ODc5LTYyZmVlOTM3MGI4ZWU0YmJkMjUxNmQ4ODQ2OTAwYzMxMTlhY2Q1MWQyOWUyNjc2ODliNjFjMDUxNWNmNzk5MDk)
[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/caddy-ui-server.svg)](https://github.com/qdm12/caddy-ui-server/commits/master)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/caddy-ui-server.svg)](https://github.com/qdm12/caddy-ui-server/graphs/contributors)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/caddy-ui-server.svg)](https://github.com/qdm12/caddy-ui-server/issues)

## Architecture

This server is stateless and does the following:

- Serve the built static files of the frontend to the user
- Act as an intermediary proxy to communicate with your Caddy server (to work around CORS)

It is written in Go in order to have a small static binary on a Docker Scratch image.

## Repository

The repository contains:

- the HTTP API server written in Go
- the web frontend app written in ReactJS, as the **ui** Git submodule

## Setup

1. Use the following command:

    ```sh
    docker run -d -p 8000:8000/tcp qmcgaw/caddy-ui
    ```

    You can also use [docker-compose.yml](https://github.com/qdm12/caddy-ui-server/blob/master/docker-compose.yml) with:

    ```sh
    docker-compose up -d
    ```

1. You can update the image with `docker pull qmcgaw/caddy-ui` or use one of [tags available](https://hub.docker.com/r/qmcgaw/caddy-ui/tags)

### Environment variables

| Environment variable | Default | Description |
| --- | --- | --- |
| `CADDY_API_ENDPOINT` | `http://localhost:2019` | Caddy server API endpoint address |
| `LOG_ENCODING` | `console` | Logging format, can be `json` or `console` |
| `LOG_LEVEL` | `info` | Logging level, can be `debug`, `info`, `warning`, `error` |
| `NODE_ID` | `-1` | Node ID for logger (`-1` to disable) |
| `LISTENING_PORT` | `8000` | Internal listening TCP port |
| `ROOT_URL` | `/` | URL path, used if behind a reverse proxy |
| `TZ` | `America/Montreal` | Timezone string |

## Development

1. Setup your environment

    <details><summary>Using VSCode and Docker (easier)</summary><p>

    1. Install [Docker](https://docs.docker.com/install/)
       - On Windows, share a drive with Docker Desktop and have the project on that partition
       - On OSX, share your project directory with Docker Desktop
    1. With [Visual Studio Code](https://code.visualstudio.com/download), install the [remote containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
    1. In Visual Studio Code, press on `F1` and select `Remote-Containers: Open Folder in Container...`
    1. Your dev environment is ready to go!... and it's running in a container :+1: So you can discard it and update it easily!

    </p></details>

    <details><summary>Locally</summary><p>

    1. Install [Go](https://golang.org/dl/), [Docker](https://www.docker.com/products/docker-desktop) and [Git](https://git-scm.com/downloads)
    1. Install Go dependencies with

        ```sh
        go mod download
        ```

    1. Install [golangci-lint](https://github.com/golangci/golangci-lint#install)
    1. You might want to use an editor such as [Visual Studio Code](https://code.visualstudio.com/download) with the [Go extension](https://code.visualstudio.com/docs/languages/go). Working settings are already in [.vscode/settings.json](https://github.com/qdm12/caddy-ui-server/master/.vscode/settings.json).

    </p></details>

1. Commands available:

    ```sh
    # Build the binary
    go build cmd/app/main.go
    # Test the code
    go test ./...
    # Lint the code
    golangci-lint run
    # Build the Docker image
    docker build -o build ui
    docker build -t qmcgaw/caddy-ui .
    # Run the container
    docker run -it --rm -p 8000:8000/tcp qmcgaw/caddy-ui
    ```

1. See [Contributing](https://github.com/qdm12/caddy-ui-server/master/.github/CONTRIBUTING.md) for more information on how to contribute to this repository.

## TODOs

- [ ] Unit tests
- [ ] Integration tests

## License

This repository is under an [MIT license](https://github.com/qdm12/caddy-ui-server/master/license) unless otherwise indicated
