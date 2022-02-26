FROM golang:1.17.3
EXPOSE 8080

RUN mkdir -p /go/src/web
WORKDIR /go/src/web

SHELL ["/bin/bash", "-c"]

# install packages
COPY . /go/src/web
RUN go mod download

RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w'

ENTRYPOINT ["/go/bin/main"]
