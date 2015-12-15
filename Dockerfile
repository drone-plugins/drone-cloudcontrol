# Docker image for the Drone cloudControl plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-cloudcontrol
#     make deps build docker

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    git \
    openssh \
    curl \
    perl && \
  rm -rf /var/cache/apk/*

ADD drone-cloudcontrol /bin/
ENTRYPOINT ["/bin/drone-cloudcontrol"]
