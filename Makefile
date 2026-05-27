.PHONY: build test validate clean all

all: test build

build:
	go run ./cmd/build

test:
	go test ./...

validate: build
	@echo "build runs schema validation; if you see 'wrote themes/veda-pro.json' the theme is valid."

clean:
	rm -f themes/veda-pro.json
