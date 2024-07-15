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

# Build api binary
RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o api ./cmd/api/*.go

FROM alpine AS app

# Setup TLS
RUN apk update && \
    apk add --no-cache ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

# Running directory has to have the same path as the build one for stacktrace mapping
WORKDIR /app

# Copy api binary
COPY --from=builder /app/api ./api

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

# Run container as non-root
RUN addgroup -S app -g 1000 && \
    adduser -u 1000 -S app -G app -h /app -s /bin/ash && \
    chown -R app:app /app

# Api
EXPOSE 1111

USER app

CMD [ "/app/api" ]
