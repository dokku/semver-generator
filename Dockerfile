FROM golang:1.22.4-alpine as builder

COPY . /go/src/github.com/dokku/semver-generator
WORKDIR /go/src/github.com/dokku/semver-generator
RUN go build

FROM alpine:3.19.1
# hadolint ignore=DL3018
COPY --from=builder /go/src/github.com/dokku/semver-generator/semver-generator /usr/local/bin/semver-generator
COPY github-entrypoint /usr/local/bin/github-entrypoint
ENTRYPOINT ["/usr/local/bin/github-entrypoint"]
LABEL org.opencontainers.image.description "A tool for generating a semver version from an existing version and a desired bump level"