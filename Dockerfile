# syntax=docker/dockerfile:1

# Comments are provided throughout this file to help you get started.
# If you need more help, visit the Dockerfile reference guide at
# https://docs.docker.com/engine/reference/builder/

################################################################################
# Create a stage for building the application.
ARG GO_VERSION=1.21
FROM golang:${GO_VERSION}-alpine AS build
WORKDIR /src

# Install build dependencies
RUN apk add --no-cache git

# Copy only the go.mod and go.sum files first
COPY go.mod go.sum ./

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x

# Copy the rest of the source code
COPY . .

# Build the application for ARM64 architecture.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
RUN --mount=type=cache,target=/go/pkg/mod/ \
    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o /bin/server .

RUN --mount=type=cache,target=/go/pkg/mod/ \
    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o /bin/bun ./cmd/bun/main.go

################################################################################
# Create a new stage for running the application that contains the minimal
# runtime dependencies for the application.
FROM --platform=linux/arm64 alpine:3 AS final

# Install any runtime dependencies that are needed to run your application.
# Leverage a cache mount to /var/cache/apk/ to speed up subsequent builds.
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update --no-cache add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

# Copy the executables from the "build" stage.
COPY --from=build /bin/server /bin/
COPY --from=build /bin/bun /bin/bun/
COPY entrypoint.sh /bin/entrypoint.sh

# Copy only necessary source files
COPY --from=build /src/go.mod /src/go.sum /src/

# Make the entrypoint script executable
RUN chmod +x /bin/entrypoint.sh

# Create a non-privileged user that the app will run under.
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# Expose the port that the application listens on.
EXPOSE 4001

# What the container should run when it is started.
ENTRYPOINT ["sh", "/bin/entrypoint.sh"]
