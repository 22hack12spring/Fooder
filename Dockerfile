FROM golang:1.18

ENV GO111MODULE=on
ENV DOCKERIZE_VERSION v0.6.1
RUN go install github.com/jwilder/dockerize@$DOCKERIZE_VERSION

# directory setting
WORKDIR /go/src/app

RUN GO111MODULE=off go get -u github.com/oxequa/realize

ENTRYPOINT dockerize -timeout 60s -wait tcp://mysql:3306 realize start