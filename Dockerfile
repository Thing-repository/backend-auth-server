# syntax=docker/dockerfile:1.2

FROM golang:1.18-alpine

RUN go version
ENV GOPATH=/

COPY ./ /thing-repository

WORKDIR /thing-repository

# build go app
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g cmd/app/main.go

RUN --mount=type=secret,id=SALT --mount=type=secret,id=TOKEN_SECRET cat /run/secrets/SALT && cat /run/secrets/TOKEN_SECRET && \
    export SALT=$(cat /run/secrets/SALT) && \
    export TOKEN_SECRET=$(cat /run/secrets/TOKEN_SECRET) && ls /run/secrets

RUN echo $TOKEN_SECRET && echo $SALT

RUN go build -o thing-repository -ldflags "-X main.tokenSecret=$TOKEN_SECRET -X main.salt=$SALT" ./cmd/app/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ENV AUTH_SERVER_ENV=OK

WORKDIR /root/

COPY --from=0 /thing-repository/thing-repository .
COPY --from=0 /thing-repository/configs/ ./configs/



CMD ["./thing-repository"]