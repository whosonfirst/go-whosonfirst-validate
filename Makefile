CWD=$(shell pwd)
GOPATH := $(CWD)

build:	rmdeps deps fmt bin

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-validate
	# cp -r *.go src/github.com/whosonfirst/go-whosonfirst-validate/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-index"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-names"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-geojson-v2"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-log"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/warning"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	cp -r src/ vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/wof-validate-index cmd/wof-validate-index.go	
