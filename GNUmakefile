TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
VERSION?=$$(git describe --abbrev=0 --tags)
BUILDDATE?=$$(date '+%Y/%m/%d %H:%M:%S %Z')
HASH?=$$(git rev-parse --verify HEAD)
GOVERSION?=$$(go version)

export GO111MODULE := on

default: build

build: fmtcheck
	go install

test: fmtcheck
	go test $(TEST) -timeout=30s -parallel=4

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w $(GOFMT_FILES)

# Currently required by tf-deploy compile
fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint:
	@echo "==> Checking source code against linters..."
	@GOGC=30 golangci-lint run ./

tools:
	GO111MODULE=on go install github.com/client9/misspell/cmd/misspell
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: build fmt fmtcheck lint tools