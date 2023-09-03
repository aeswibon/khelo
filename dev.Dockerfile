# syntax=docker/dockerfile:1
ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS base

FROM base AS builder

WORKDIR /usr/src/app

# install dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# install air
RUN go install github.com/cosmtrek/air@latest

# copy the source code
COPY . .

ENTRYPOINT ["air", "-c", ".air.toml"]