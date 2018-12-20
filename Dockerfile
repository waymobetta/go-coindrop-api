FROM golang:latest

RUN mkdir /app

RUN git config --global url."https://825d179e916f393a4abfc86a20828facda0169d2:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/waymobetta/wmb
RUN go get github.com/jzelinskie/geddit
RUN go get golang.org/x/crypto/bcrypt

ADD . /app

WORKDIR /app/github.com/waymobetta/go-coindrop-api/cmd

RUN go build -o ../bin/go-coindrop-api

EXPOSE 8000

CMD ["/app/bin/go-coindrop-api"]
