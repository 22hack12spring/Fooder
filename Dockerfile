FROM golang:1.18

ENV GO111MODULE=on

# directory setting
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app

RUN GO111MODULE=off go get -u github.com/oxequa/realize

CMD ["realize", "start"]