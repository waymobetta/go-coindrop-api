FROM golang:latest

# expose default port
EXPOSE 5000

# set environemnt variable path
ENV PATH /go/bin:$PATH

# cd into directory
WORKDIR /go/src/github.com/waymobetta/go-coindrop-api

# allow private repo pull
RUN git config --global url."https://825d179e916f393a4abfc86a20828facda0169d2:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# install dependencies
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/waymobetta/wmb
RUN go get github.com/jzelinskie/geddit
RUN go get golang.org/x/crypto/bcrypt

# copy local package files to container workspace
ADD . /go/src/github.com/waymobetta/go-coindrop-api

# install program
RUN go install github.com/waymobetta/go-coindrop-api/cmd

# start app
CMD ["cmd"]
