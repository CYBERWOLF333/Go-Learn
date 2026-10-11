.PHONY: all build test clean demo run

BINARY_NAME=bin/gopherrecon

all: test build

build:
	@mkdir -p bin
	go build -o $(BINARY_NAME) ./cmd/gopherrecon

test:
	go test -v ./...

demo: build
	./$(BINARY_NAME) -demo

clean:
	rm -rf bin demo_recon_report.*
