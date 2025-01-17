FROM golang:alpine3.18 as builder
WORKDIR /golang-api-json-in-docker
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" .
FROM busybox
WORKDIR /golang-api-json-in-docker
COPY --from=builder /golang-api-json-in-docker /usr/bin/
ENTRYPOINT ["golang-api-json-in-docker"]