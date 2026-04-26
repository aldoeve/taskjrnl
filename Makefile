run:
	clip <(echo go run ./cmd/taskjrnl/ add "Hello")

help:
	go run ./cmd/taskjrnl/ help

test:
	go test ./...

verbose:
	go test ./... -v

clean:
	rm -f ./internal/store/TJ.db