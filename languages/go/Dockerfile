FROM golang:1.13-alpine as builder
WORKDIR /usr/src/app

ENV SERVER_PORT=50101

# Enable private repositories
RUN apk add --no-cache --virtual .build-deps g++ git
ENV GOPRIVATE=github.com/upnid/*
RUN git config --global url."https://59eb5e63e03c411c701676e678d0cc12d6fbbc1d:x-oauth-basic@github.com/upnid".insteadOf "https://github.com/upnid"

# Build
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN go build -o main .

CMD ["./main"]
EXPOSE $SERVER_PORT
