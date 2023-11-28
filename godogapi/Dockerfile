FROM golang:1.21.4

ENV GIN_MODE=release
ENV API_PORT=3004

WORKDIR /api

COPY . /api

RUN go build -o ./api cmd/main.go

EXPOSE $PORT

ENTRYPOINT ["./api"]