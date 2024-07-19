# Make sure we always get a statically linked binary that does not require any additional software from the OS
export CGO_ENABLED := 0

# List of all .go files in the project. Used by Make to determine if the project needs to be rebuilt.
GOFILES := $(shell find . -type f -name '*.go' -not -name '*_gen.go')
# These go directly to the `go build` command. It's used for passing multiple options to `-ldflags` which is currently
# not supported via the standard `GOFLAGS` env variable - that only supports a single option.
# See: https://github.com/golang/go/issues/38522
GO_FLAGS := $(GO_FLAGS)

# Default target
all: compile

### File targets

dist:
	mkdir -p $@

go.sum: go.mod
	go mod tidy
	touch go.sum

dist/gqlgen.sentinel: dist graph/*.graphqls
	go run graph/gqlgen.go
	touch $@

dist/helloserver: go.sum dist/gqlgen.sentinel $(GOFILES)
	go build -v -o $@ $(GO_FLAGS)


### Human-friendly targets

generate: dist/gqlgen.sentinel

compile: dist/helloserver

lint: generate
	go fmt ./...
	go vet ./...
	./tools/golangci-lint run -v

test: generate
	go test -v ./...

container:
	docker build -t helloserver:latest .

distclean:
	rm -rf dist

clean: distclean
	rm -rf .cache
	git clean -Xf .
