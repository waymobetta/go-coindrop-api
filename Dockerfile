FROM golang:latest

WORKDIR /go/src/github.com/waymobetta/go-coindrop-api

RUN git config --global url."https://825d179e916f393a4abfc86a20828facda0169d2:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/waymobetta/wmb
RUN go get github.com/jzelinskie/geddit
RUN go get golang.org/x/crypto/bcrypt

ADD . /go/src/github.com/waymobetta/go-coindrop-api

RUN go install github.com/waymobetta/go-coindrop-api/cmd

EXPOSE 5000

CMD ["cmd"]
