SHELL    = /bin/bash
GOPATH   = $(shell pwd)
PLATFORM = $(shell uname)

GO        = GOPATH=$(GOPATH) /usr/local/go/bin/go

DEPS      = github.com/valyala/fasthttp

ifeq ($(PLATFORM),Darwin)
BROWSER = open
else ifeq ($(PLATFORM),Linux)
BROWSER = x-www-browser
endif

all: clean $(DIRS) goget.go
	$(MAKE) fast

fast:
	$(GO) build -o goget goget.go

nuke:
	-rm -rf src/github.com src/golang.org src/gopkg.in pkg
	-rm -rf pkg/
	-rm -rf bin/

clean:
	@-rm -f goget

update:
	for dep in $(DEPS) ; do $(GO) get -u $@ ; done
