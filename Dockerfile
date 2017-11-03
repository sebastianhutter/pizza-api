#
# compile and run a small http to https redirect service
#

FROM alpine:3.6
MAINTAINER Sebastian Hutter <mail@sebastian-hutter.ch>

WORKDIR /
COPY main.go /main.go
COPY Makefile /Makefile

# compile http-redirect
RUN apk add --no-cache --update libc-dev make git go tini ca-certificates \
  && make compile \
  && apk del --purge libc-dev go git \
  && mv /bin/pizza-api / \
  && rm /main.go \
  && rm /Makefile

EXPOSE 80
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/pizza-api"]