# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY ./binance ./binance
COPY ./utils ./utils

RUN CGO_ENABLED=0 GOOS=linux go build -o /binance-price

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /binance-price /binance-price

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/binance-price"]