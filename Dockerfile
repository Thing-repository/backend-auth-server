# syntax=docker/dockerfile:1.2

FROM golang:1.19-alpine

RUN go version
ENV GOPATH=/

COPY ./ /thing-repository

WORKDIR /thing-repository

# build go app
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g cmd/app/main.go

RUN go build -o thing-repository ./cmd/app/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates


WORKDIR /root/

COPY --from=0 /thing-repository/thing-repository .
COPY --from=0 /thing-repository/configs/ ./configs/

CMD ["./thing-repository"]