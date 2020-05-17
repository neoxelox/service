FROM golang:1.14.2 AS builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

# Download dependencies before build in order to cache them
COPY go.mod go.sum /src/

RUN go mod download

COPY . /src

RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o mst ./cmd/mst/main.go

FROM alpine AS app

RUN apk add ca-certificates

WORKDIR /app

# Copy source files for stacktrace mapping
COPY cmd /src/cmd
COPY internal /src/internal
COPY pkg /src/pkg

COPY assets /assets
COPY --from=builder /src/mst /app

# App
EXPOSE 8000

CMD [ "/app/mst" ]