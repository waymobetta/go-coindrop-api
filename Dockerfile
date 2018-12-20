FROM golang:latest

EXPOSE 9000

ENV PATH /go/bin:$PATH

WORKDIR /go/src/github.com/waymobetta/go-coindrop-api

RUN MKDIR ~/.ssh
RUN TOUCH ~/.ssh/known_hosts
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

RUN git config --global url."https://825d179e916f393a4abfc86a20828facda0169d2:x-oauth-basic@github.com/".insteadOf "https://github.com/"

ADD . /go/src/github.com/waymobetta/go-coindrop-api

RUN go install github.com/waymobetta/go-coindrop-api

CMD ["make", "start"]

