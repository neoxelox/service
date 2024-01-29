FROM golang:1.21-alpine

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Setup air hot reloader
RUN wget https://github.com/cosmtrek/air/releases/download/v1.49.0/air_1.49.0_linux_amd64 -O air && \
    chmod +x air && \
    mv ./air /bin/air

# Api
EXPOSE 1111

# Gilk
EXPOSE 1113

CMD [ "air", "-c", "./envs/dev/api.air.toml" ]
