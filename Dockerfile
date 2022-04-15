# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o ./mockapp

FROM alpine:latest

RUN apk update && apk upgrade && apk add curl jq

WORKDIR /app

COPY --from=build /app/mockapp .
COPY config.yaml ./
COPY public public
