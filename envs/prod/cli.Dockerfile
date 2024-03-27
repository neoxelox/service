FROM golang:1.22 AS builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

# Build directory has to have the same path as the running one for stacktrace mapping
WORKDIR /app

# Download dependencies before build in order to cache them
COPY go.mod go.sum ./
RUN go mod download

# Copy source files for compiling
COPY cmd ./cmd
COPY pkg ./pkg

# Build cli binary
RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o cli ./cmd/cli/*.go

FROM alpine AS app

# Setup TLS
RUN apk update && \
    apk add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/*

# Running directory has to have the same path as the build one for stacktrace mapping
WORKDIR /app

# Copy cli binary
COPY --from=builder /app/cli ./cli

# Copy source files for stacktrace mapping
COPY cmd ./cmd
COPY pkg ./pkg

# Copy other resources
COPY assets ./assets
COPY locales ./locales
COPY migrations ./migrations
COPY templates ./templates

# Create files directory
RUN mkdir -p ./files

# Set CLI binary on PATH
ENV PATH="$PATH:/app"

# Run container as non-root
RUN addgroup -S app -g 1000 && \
    adduser -u 1000 -S app -G app -h /app -s /bin/ash && \
    chown -R app:app /app

USER app

ENTRYPOINT [ "/app/cli" ]
