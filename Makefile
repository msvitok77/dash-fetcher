BINARY_NAME=dash-fetcher
LDFLAGS=-s -w

all: build

build:
	go build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) cmd/$(BINARY_NAME)/main.go

start-file-server:
	cd cmd/file-server; \
	./run.sh

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)

.PHONY: all build test clean
