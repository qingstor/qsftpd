SHELL := /bin/bash

VERSION=$(shell cat version.go | grep "Version\ =" | sed -e s/^.*\ //g | sed -e s/\"//g)
DIRS_WITHOUT_VENDOR=$(shell ls -d */ | grep -vE "vendor")
PKGS_WITHOUT_VENDOR=$(shell go list ./... | grep -v "/vendor/")

.PHONY: help
help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  all               to check, build, test and release"
	@echo "  check             to vet and lint"
	@echo "  build             to create bin directory and build"
	@echo "  test              to run test"
	@echo "  run               to run qsftp locally"
	@echo "  integration-test  to run integration test"
	@echo "  release           to build and release"
	@echo "  clean             to clean build and test files"

.PHONY: all
all: check build test release clean

.PHONY: check
check: vet lint

.PHONY: vet
vet:
	@echo "go tool vet, on qsftp packages"
	@go tool vet -all ${DIRS_WITHOUT_VENDOR}
	@echo "ok"

.PHONY: lint
lint:
	@echo "golint, on qsftp packages"
	@lint=$$(for pkg in ${PKGS_WITHOUT_VENDOR}; do golint $${pkg}; done); \
	 if [[ -n $${lint} ]]; then echo "$${lint}"; exit 1; fi
	@echo "ok"

.PHONY: build
build:
	@echo "build qsftp"
	mkdir -p ./bin
	go build -o ./bin/qsftp .
	@echo "ok"

.PHONY: test
test:
	@echo "run test"
	@go test -v ${PKGS_WITHOUT_VENDOR}
	@echo "ok"

.PHONY: run
run: build
	@echo "run qsftp"
	./bin/qsftp -c qsftp.yaml
	@echo "ok"

.PHONY: integration-test
integration-test:
	@if [[ ! -f "$$(which bats)" ]]; then \
		echo "ERROR: Command \"bats\" not found."; \
	fi
	@echo "run integration test"
	bats test
	@echo "ok"

.PHONY: release
release:
	@echo "release qsftp"
	mkdir -p ./release
	@echo "for Linux"
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux/qsftp .
	mkdir -p ./release
	tar -C ./bin/linux/ -czf ./release/qsftp-v${VERSION}-linux_amd64.tar.gz qsftp
	@echo "for macOS"
	mkdir -p ./bin/linux
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/qsftp .
	tar -C ./bin/darwin/ -czf ./release/qsftp-v${VERSION}-darwin_amd64.tar.gz qsftp
	@echo "for Windows"
	mkdir -p ./bin/windows
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows/qsftp.exe .
	tar -C ./bin/windows/ -czf ./release/qsftp-v${VERSION}-windows_amd64.tar.gz qsftp.exe
	@echo "ok"

.PHONY: clean
clean:
	rm -rf ./bin
	rm -rf ./release
	rm -rf ./coverage
