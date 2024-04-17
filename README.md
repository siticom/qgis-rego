# QGIS-rego

This repository contains an implementation of a QGIS Plugin Repository using Go. Our goal is to provide a feature-rich and easy-to-deploy solution.

## Run with Docker Compose

```sh
# run docker-compose
docker-compose up -d
# open on http://localhost:8080/plugins/plugins.xml
```

## Configuration

Configuration for the application is handled by the `Configuration` struct in [config/env.go](config/env.go). The following environment variables can be set:

- `MINIMUM_QGIS_VERSION`: The minimum version of QGIS that the plugins support.
- `PLUGIN_DIR`: The directory where the plugins are stored.
- `PORT`: The port on which the application listens.
- `SHOW_ICONS`: A boolean indicating whether to show icons (includes Base64 encoded icons in the response)

See [.env.example](.env.example) for an example configuration.

## Development

Requires [Go](https://go.dev/) to be installed.

```sh
# optional: copy env file
cp .env.example .env

# run program
go run .
# open on http://localhost:8080/plugins/plugins.xml

# build
go build
# run binary with
./qgis-repository
```

## Similar Projects

- [PHP variant of Michel Stuyts](https://gitlab.com/GIS-projects/phpQGISrepository)
- [Official Django Plugin repository](https://github.com/qgis/QGIS-Django/tree/master/qgis-app/plugins)
