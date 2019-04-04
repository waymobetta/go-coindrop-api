FROM golang:1.12rc1-alpine3.9 AS build

RUN apk --no-cache add ca-certificates git build-base
COPY . /go/src/github.com/waymobetta/go-coindrop-api
RUN rm -r /go/src/github.com/waymobetta/go-coindrop-api/vendor/github.com/ethereum/go-ethereum
RUN go get -u github.com/ethereum/go-ethereum
WORKDIR /go/src/github.com/waymobetta/go-coindrop-api
RUN apk del git

RUN GOOS=linux go build -a -installsuffix cgo -o coindrop cmd/coindrop/main.go

FROM alpine:latest

WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/waymobetta/go-coindrop-api/coindrop .

# expose default port
EXPOSE 5000

# start app
CMD ["./coindrop"]