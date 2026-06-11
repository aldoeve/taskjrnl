BIN_DIR := $(shell go env GOPATH)/bin

run:
	go run ./cmd/taskjrnl/ add "Hello"

help:
	go run ./cmd/taskjrnl/ help

test:
	go test ./...

verbose:
	go test ./... -v

coverage:
	go test ./... -coverprofile=coverage.out -count=1
	go tool cover -html=coverage.out -o coverage.html
	xdg-open coverage.html

install:
	go install ./cmd/taskjrnl/

symlink:
	ln -sf $(BIN_DIR)/taskjrnl $(BIN_DIR)/task

asm:
	@uuid=$$(uuidgen); \
	go build -gcflags="-S" ./... 2> asm$${uuid}.txt;

clean:
	rm -f coverage.out
	rm -f coverage.html
	