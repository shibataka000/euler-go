BUILD_TARGET = github.com/shibataka000/euler-go/cmd
FMT_TARGET = $(shell find . -type f -name "*.go")
LINT_TARGET = $(shell go list ./...)
TEST_TARGET = ./...
BUILD_OUTPUT = solver.out
BUILD_VERSION = $(shell git describe --tags)
BUILD_FLAGS = -ldflags "-X main.version=$(BUILD_VERSION)"

default: build

setup:
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports

checkfmt:
	test ! -n "$(shell goimports -l $(FMT_TARGET))"

lint:
	go vet $(LINT_TARGET)
	golint $(LINT_TARGET)

test:
	go test $(TEST_TARGET)

build:
	go build -o $(BUILD_OUTPUT) $(BUILD_FLAGS) $(BUILD_TARGET)

ci: checkfmt lint test build

.PHONY: default setup checkfmt lint test build ci
