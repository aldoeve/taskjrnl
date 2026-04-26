run:
	clip <(echo go run ./cmd/taskjrnl/ add "Hello")

help:
	go run ./cmd/taskjrnl/ help

what:
	@echo "run, help, test, verbose, coverage, clean"

test:
	go test ./...

verbose:
	go test ./... -v

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	xdg-open coverage.html

clean:
	rm -f ./internal/store/TJ.db
	rm -f coverage.out
	rm -f coverage.html
	