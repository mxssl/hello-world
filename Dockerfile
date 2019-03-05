FROM golang:1.12.0-alpine3.9 as builder

ENV GO111MODULE=on

WORKDIR /go/src/github.com/mxssl/hello-world
COPY . .

# Компилируем бинарник
RUN apk add --no-cache ca-certificates git
RUN CGO_ENABLED=0 \
  GOOS=`go env GOHOSTOS` \
  GOARCH=`go env GOHOSTARCH` \
  go build -o app

# Копируем скомпилированный бинарник в чистый образ Alpine Linux
FROM alpine:3.9
WORKDIR /
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/mxssl/hello-world .
RUN chmod +x app
CMD ["./app"]
