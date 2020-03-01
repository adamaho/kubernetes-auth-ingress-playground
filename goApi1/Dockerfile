
FROM golang:latest AS gobuilder
RUN go version
WORKDIR /go/src/github.com/rtkaaho/api1
COPY . .
RUN set -x && \
  go get github.com/golang/dep/cmd/dep && \
  dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o api1 .

FROM alpine:latest
WORKDIR /root/
COPY --from=gobuilder /go/src/github.com/rtkaaho/api1/ .
CMD ["./api1"]
EXPOSE 8080