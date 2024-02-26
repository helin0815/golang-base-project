# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: all build clean cover vendor fmt fmtcheck generate test testacc testrace tools website

TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
TESTARGS?=-gcflags=-l
VERSION := $(shell git describe --tags)
LDFLAGS="-X gitlabee.chehejia.com/serverless/faasapi/bootstrap.Version=$(VERSION)"
default: test

build:
	@echo "Building version: $(VERSION)"
	@echo $(shell git describe --tags --abbrev=0) > ".version"
	go build -v -a -installsuffix cgo -ldflags $(LDFLAGS) -o output/main.linux ./main.go

coverhtml: test
	go tool cover -html=cover.out

cover: test
	go tool cover -func=cover.out

test:
	test -s $(GOPATH)/bin/gotestsum || go install gotest.tools/gotestsum@latest
	$(GOPATH)/bin/gotestsum --junitfile report.xml --jsonfile report.json --format testname -- -coverpkg=$(shell go list)/... -coverprofile=cover.out ./...

# testrace runs the race checker
testrace: generate
	TF_ACC= go test -race $(TEST) $(TESTARGS)

# generate runs `go generate` to build the dynamically generated source files
generate: generate-tools
	go run hack/dalgen/main.go
	go generate ./...

run:
	go run main.go

clean:
	rm -fr bin/*

generate-tools:
	test -s yq || go install github.com/mikefarah/yq/v4@latest