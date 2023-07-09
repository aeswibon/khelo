# syntax=docker/dockerfile:1
ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS base

FROM base AS builder

# install dependencies
RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,src=go.mod,target=go.mod \
  --mount=type=bind,src=go.sum,target=go.sum \
  go mod download

FROM builder AS dev

WORKDIR /app

ENTRYPOINT ["air", "-c", ".air.toml"]