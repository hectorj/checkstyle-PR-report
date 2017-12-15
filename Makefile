SHELL=/usr/bin/env bash -euo pipefail

all: lint-results.checkstyle.xml gotest.report.txt report.html

build:
	mkdir ./build

build/packr: build
	GOBIN="$$(pwd)/build/" go get -v github.com/gobuffalo/packr/packr

build/gometalinter: build
	GOBIN="$$(pwd)/build/" go get -v github.com/alecthomas/gometalinter && ./build/gometalinter --config=gometalinter.conf.json --install

build/ir-blaster: build build/packr
	GOBIN="$$(pwd)/build/" ./build/packr install ./ir-blaster

build/lint-results.checkstyle.xml: build build/gometalinter
	./build/gometalinter --disable-all --config=gometalinter.conf.json --checkstyle ./... | tee ./build/lint-results.checkstyle.xml

build/gotest.report.txt: build
	go test -v ./... 2>&1 | tee ./build/gotest.report.txt

report-to-github: build/ir-blaster build/lint-results.checkstyle.xml build/gotest.report.txt build/ir-blaster
	./build/ir-blaster github --gotest="./build/gotest.report.txt" --checkstyle="./build/lint-results.checkstyle.xml" --github-repo-owner=hectorj --github-repo-name=ir-blaster --github-pr-id=$${TRAVIS_PULL_REQUEST} --github-oauth-token=$${GITHUB_TOKEN}

build/report.html: build build/ir-blaster build/lint-results.checkstyle.xml build/gotest.report.txt build/ir-blaster
	./build/ir-blaster htmlfile --gotest="./build/gotest.report.txt" --checkstyle="./build/lint-results.checkstyle.xml" --output="./build/report.html"

.PHONY: all build/ir-blaster build/gotest.report.txt build/lint-results.checkstyle.xml build/report.html report-to-github #TODO: define dependencies and stop using PHONY
