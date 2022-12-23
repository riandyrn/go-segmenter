.PHONY: *

test:
	go test ./... -count=1 -timeout=30s

build:
	go build