FROM alpine

RUN apk --no-cache add curl

COPY bin/amd64/httpserver /httpserver

ENTRYPOINT /httpserver