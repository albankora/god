FROM golang:alpine3.8

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates git bash

ENV APP_ENV dev
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY ./build/dev/entrypoint.sh /usr/local/bin/entrypoint.sh

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /usr/local/bin/entrypoint.sh

RUN go get github.com/oxequa/realize

WORKDIR $GOPATH

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

CMD ["realize", "start"]