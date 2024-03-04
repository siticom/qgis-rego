# QGIS-rego

This repository contains an implementation of a QGIS Plugin Repository using Go. Our goal is to provide a feature-rich and easy-to-deploy solution.

## Prerequisites

- Docker
- docker-compose

## Getting Started

1. Clone the repository.
2. Copy the `.env.example` file to `.env` and adjust the values to your needs.
3. Run `docker-compose up -d` to start the repository.

## Building

The project uses a multi-stage Docker build. The `Dockerfile` first creates a build stage container where it downloads the dependencies and builds the Go application. Then it creates a release stage container where it copies the built application and the static files. The application listens on port 8080.

## Configuration

Configuration for the application is handled by the `Configuration` struct in [config/env.go](config/env.go). The following environment variables can be set:

- `MINIMUM_QGIS_VERSION`: The minimum version of QGIS that the plugins support.
- `PLUGIN_DIR`: The directory where the plugins are stored.
- `PORT`: The port on which the application listens.
- `SHOW_ICONS`: A boolean indicating whether to show icons (includes Base64 encoded icons in the response)

See [.env.example](.env.example) for an example configuration.

## Similar Projects

- [PHP variant of Michel Stuyts](https://gitlab.com/GIS-projects/phpQGISrepository)
- [Official Django Plugin repository ](https://github.com/qgis/QGIS-Django/tree/master/qgis-app/plugins)