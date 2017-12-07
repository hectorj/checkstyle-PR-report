all: _testdata/ab0x.go static/ab0x.go test-results.junit.xml lint-results.checkstyle.xml report.html

_testdata/ab0x.go: _testdata/b0x.toml
	cd _testdata && rm -f ab0x.go && fileb0x b0x.toml

static/ab0x.go: static/b0x.toml
	cd static && rm -f ab0x.go && fileb0x b0x.toml

lint-results.checkstyle.xml:
	#gometalinter --disable-all --config=gometalinter.conf.json --checkstyle ./... > lint-results.checkstyle.xml || true

report.html: lint-results.checkstyle.xml
	go test -v -cover ./... 2>&1 | go run ./ir-blaster-go/main.go --gotest="stdin://" --checkstyle="./lint-results.checkstyle.xml" --out="./report.html"

.PHONY: all _testdata/ab0x.go static/ab0x.go test-results.junit.xml lint-results.checkstyle.xml report.html #TODO: define dependencies and stop using PHONY
