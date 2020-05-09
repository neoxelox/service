FROM golang:1.14.2 AS builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . /build

RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o mst ./cmd/mst/main.go

FROM alpine AS app

RUN apk add ca-certificates

WORKDIR /app

COPY ./assets /assets
COPY --from=builder /build/mst /app

# App
EXPOSE 8000

CMD [ "/app/mst" ]