FROM golang:1.10-stretch as base-deps
WORKDIR /go/src/github.com/ovrclk/akash
COPY . /go/src/github.com/ovrclk/akash
RUN set -eu; \
    curl https://glide.sh/get | sh ; \
    glide --home /glide install -v ; \
    make bins
