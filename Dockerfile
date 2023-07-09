# syntax=docker/dockerfile:1
ARG GO_VERSION=1.20
ARG ALPINE_VERSION=3.18.2

# base stage
FROM golang:${GO_VERSION}-alpine AS base

# builder stage
FROM base AS builder

WORKDIR /app

# install dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy the source code
COPY . .

# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/main.go

# prod stage
FROM alpine:${ALPINE_VERSION} AS prod

# copy the binary from builder
COPY --from=builder /go/bin/app /app

# copy the env file
COPY .env.example .env

# for go runtime
EXPOSE 8080

# for db connection
EXPOSE 27017

ENTRYPOINT ["/app"]