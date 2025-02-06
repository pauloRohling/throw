PACKAGES = ./...

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## install: install all dependencies
.PHONY: install
install:
	go mod tidy -e


## test: run all tests
.PHONY: test
test:
	go test -race -failfast -buildvcs $(PACKAGES)

## test/c: run all tests and display coverage
.PHONY: test/c
test/c:
	go test -v -race -buildvcs -coverprofile=./tmp/coverage.out $(PACKAGES)
	go tool cover -html=./tmp/coverage.out

## test/v: run all tests in verbose mode
.PHONY: test/v
test/v:
	go test -v -race -failfast -buildvcs $(PACKAGES)
	
## publish: update pkg.go.dev - tag=?
.PHONY: publish
publish:
	GOPROXY=proxy.golang.org go list -m github.com/pauloRohling/throw@$(tag)
