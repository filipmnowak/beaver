.PHONY: clean
clean:
	go clean -cache

.PHONY: build
build: clean
	go build

.PHONY: test
test:
	go test -v ./...
