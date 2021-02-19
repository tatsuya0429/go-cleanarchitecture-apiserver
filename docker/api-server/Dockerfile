FROM golang:1.16-alpine3.13 AS build

ENV GO111MODULE=on
WORKDIR /go/src/github.com/ulwlu/go-rest
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src/. ./
RUN go build

FROM alpine:3.13
COPY --from=build /go/src/github.com/ulwlu/go-rest/go-rest /usr/local/bin/

EXPOSE 8080
CMD ["go-rest"]
