FROM golang:1.16.5-buster

ENV GOPATH /.go

RUN apt-get -y update && \
    apt-get -y install git

RUN go get -u github.com/cosmtrek/air

CMD ["go", "run", "main.go"]