FROM golang:1.13-alpine as builder

ENV SERVER_HOST=0.0.0.0
ENV SERVER_PORT={{port}}
ENV TZ=UTC

WORKDIR /usr/src/app

# Enable private repositories
RUN apk add --no-cache git
ENV GOPRIVATE=github.com/upnid/*
RUN git config --global url."https://c8bec607a68f8ab97c50bc12e603f936c7db1fb6:x-oauth-basic@github.com/upnid".insteadOf "https://github.com/upnid"

COPY go.mod go.sum ./
RUN go mod download
RUN go env
COPY . . 

RUN go build .

CMD ["./{{name}}"]

EXPOSE $SERVER_PORT

