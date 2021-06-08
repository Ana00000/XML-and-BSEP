FROM golang:1.15-alpine3.13 AS build
WORKDIR /go/src/app
RUN apk update \
    && apk add socat \
    && apk add build-base
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./bin/tag-service

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8090
ENTRYPOINT /go/bin/tag-service