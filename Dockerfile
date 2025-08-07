FROM golang:1.24.6-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /qgis-rego

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /qgis-rego /qgis-rego
COPY static ./static

EXPOSE 8080

ENTRYPOINT [ "/qgis-rego" ]
