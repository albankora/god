FROM golang:alpine as builder
# Install git
RUN apk update && apk add git
# Create appuser
RUN adduser -D -g '' appuser
COPY ./src /src
WORKDIR /src
#get dependancies
RUN go get -d -v

#build the binary with debug information removed
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -a -installsuffix cgo -o main

# start from scratch
FROM scratch

# Set base enviroment variables
ENV APP_ENV prod

# Copy our static executable and pass
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /src/main /app/main
USER appuser

ENTRYPOINT ["/app/main"]