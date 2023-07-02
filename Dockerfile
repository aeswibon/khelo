# syntax=docker/dockerfile:1
ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS base

FROM base AS builder

WORKDIR /app

# install dependencies
COPY go.mod go.sum ./
RUN go mod download

FROM builder AS build-server

WORKDIR /app

# copy the packages
COPY --from=builder /go/src /go/src

# copy the source code
COPY . .

# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

FROM base AS final

# copy the binary from builder
COPY --from=build-server /server /app

EXPOSE 8080

ENTRYPOINT ["/app"]