# syntax=docker/dockerfile:1
ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS base

FROM base AS builder

# install dependencies
RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,src=go.mod,target=go.mod \
  --mount=type=bind,src=go.sum,target=go.sum \
  go mod download


FROM builder AS build-server

# build the binary
RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,src=.,target=. \
  CGO_ENABLED=0 GOOS=linux go build -o /app

FROM base AS final

# copy the binary from builder
COPY --from=builder /app /app

EXPOSE 8080

ENTRYPOINT ["/app"]