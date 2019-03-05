BINARY_NAME=app

.PHONY: all build clean lint test docker-build docker-release docker-test up run

all: dep build

build:
	go build -o ${BINARY_NAME} -v
	chmod +x app

clean:
	rm -f ${BINARY_NAME}

lint:
	golangci-lint run

test:
	go test -v ./...

docker-build:
	docker-compose build

docker-release:
	docker-compose build
	docker push mxssl/hello-world

docker-test:
	docker run \
	-e GO111MODULE=on \
	-e CGO_ENABLED=0 \
	-i \
	--rm \
	-v $(shell pwd):/go/src/github.com/mxssl/hello-world \
	-w /go/src/github.com/mxssl/hello-world \
	golang:1.12.0-alpine3.9 \
	sh -c "apk add --no-cache ca-certificates git; go test -v ./..."

up:
	docker-compose pull
	docker-compose up

run:
	go run main.go

kubernetes:
	kubectl apply -f kube/

dgoss:
	dgoss run mxssl/hello-world
