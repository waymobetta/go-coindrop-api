FROM golang:1.12rc1-alpine3.9 AS build

RUN apk --no-cache add ca-certificates
COPY . /go/src/github.com/waymobetta/go-coindrop-api
WORKDIR /go/src/github.com/waymobetta/go-coindrop-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o coindrop cmd/coindrop/main.go

FROM scratch

WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/waymobetta/go-coindrop-api/coindrop .

# expose default port
EXPOSE 5000

# start app
CMD ["./coindrop"]
