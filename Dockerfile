FROM ubuntu:20.04 as base
RUN mkdir -p /go-control-plane && chown -R nobody:nogroup /go-control-plane

COPY output/go-control-plane /usr/local/bin/go-control-plane

USER nobody
WORKDIR /go-control-plane
ENTRYPOINT ["go-control-plane"]
