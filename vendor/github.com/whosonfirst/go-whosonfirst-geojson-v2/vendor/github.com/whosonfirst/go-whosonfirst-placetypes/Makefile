CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps deps fmt bin

self:   prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-placetypes/filter
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-placetypes/placetypes
	cp *.go src/github.com/whosonfirst/go-whosonfirst-placetypes/
	cp *.go src/github.com/whosonfirst/go-whosonfirst-placetypes/
	cp filter/*.go src/github.com/whosonfirst/go-whosonfirst-placetypes/filter/
	cp placetypes/*.go src/github.com/whosonfirst/go-whosonfirst-placetypes/placetypes/
	cp -r vendor/* src

deps:   
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/warning"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-cli"

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt placetypes/*.go
	go fmt filter/*.go

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/wof-placetype-ancestors cmd/wof-placetype-ancestors.go
	@GOPATH=$(GOPATH) go build -o bin/wof-placetype-children cmd/wof-placetype-children.go
	@GOPATH=$(GOPATH) go build -o bin/wof-placetype-descendants cmd/wof-placetype-descendants.go

test:	self
	@GOPATH=$(GOPATH) go run cmd/test.go

spec:
	@GOPATH=$(GOPATH) go run cmd/mk-spec.go > placetypes/spec.go
