# We use Alpine as a base image for development and testing. Here you can override
# the version of golang-alpine to use.
ARG ALPINE_VERSION

# Define or override the Golang version used within local testing.
ARG GO_VERSION

# This will fetch the latest golang-alpine image from docker hub. Go version and
# Alpine versions are taken as --build-args during the build process.
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

# Specify the Goreleaser version to use for local testing.
ARG GORELEASER_VERSION

# Specify the local golangci-lint verion to use for local testing.
ARG GOLANGCILINT_VERSION

# Installs additional packages that may be used for local testing.
RUN apk update && apk --no-cache add yamllint make git gcc libc-dev bash gcc-go

# Install goreleaser from the specified tag in Github.
RUN wget https://github.com/goreleaser/goreleaser/releases/download/v${GORELEASER_VERSION}/goreleaser_${GORELEASER_VERSION}_x86_64.apk
RUN apk add --allow-untrusted goreleaser_${GORELEASER_VERSION}_x86_64.apk

# Install golangci-lint, also specifying the version tag to be used.
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCILINT_VERSION}

# Work around new security warning introduced in Git 2.34.2
RUN git config --global --add safe.directory /workdir

# Compose will map the current development directory on the users machine to the `/source` directory
# within the container. We set the working directory to be the mapped volume.
WORKDIR /workdir
