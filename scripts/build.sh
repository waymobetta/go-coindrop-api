#!/bin/bash

go get github.com/gorilla/mux
go get github.com/lib/pq
go build -o bin/go-coindrop-api cmd/start.go