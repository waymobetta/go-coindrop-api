FROM golang:1.11.5 AS build

COPY . /go/src/github.com/waymobetta/go-coindrop-api
WORKDIR /go/src/github.com/waymobetta/go-coindrop-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o coindrop cmd/coindrop/main.go

FROM scratch

WORKDIR /
COPY --from=build /go/src/github.com/waymobetta/go-coindrop-api/coindrop .

# expose default port
EXPOSE 5000

# start app
CMD ["./coindrop"]
