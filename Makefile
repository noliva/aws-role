.PHONY: build test test-junit install

build:
	@mkdir -p bin
	$(eval VERSION=$(shell awk '/^## ([0-9]\.[0-9]\.[0-9])/{print $$2}' CHANGELOG.md | head -n 1))
	@go build -v -i -ldflags "-X github.com/ryanfrench/aws-role/cmd.Version=${VERSION}" -o ./bin/aws-role-${VERSION}

test:
	go test -v ./...

test-junit:
	go get -u github.com/jstemmer/go-junit-report
	go test -v ./... 2>&1 | go-junit-report > report.xml

install: test
	go install
