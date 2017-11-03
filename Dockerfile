#
# compile and run a small http to https redirect service
#

FROM alpine:3.6
MAINTAINER Sebastian Hutter <mail@sebastian-hutter.ch>

WORKDIR /
COPY main.go /build/main.go
COPY Makefile /build/Makefile

# compile http-redirect
RUN apk add --no-cache --update libc-dev make git go tini ca-certificates \
  && cd /build \
  && make compile \
  && apk del --purge libc-dev go git \
  && mv /build/bin/pizza-api / \
  && rm -rf /build

EXPOSE 80
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/pizza-api"]