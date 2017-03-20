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

DEPDIR  = $(patsubst %, src/%, $(DEPS))

SRCDIRS	= $(shell find src/goget -name '*.go' -print | while read file ; do echo `dirname $$file` ; done | sort -u)
DIRS = $(patsubst src/%, %, $(SRCDIRS))

all: clean $(DEPDIR) goget.go
	$(MAKE) fast

fast:
	$(GO) build -tags "prod" -o goget goget.go

nuke:
	-rm -rf src/github.com src/golang.org src/gopkg.in pkg
	-rm -rf pkg/
	-rm -rf bin/

clean:
	@-$(GO) clean $(DIRS)
	@-rm -f goget

$(DEPDIR):
	$(GO) get -v $(patsubst src/%, %, $@)

update:
	for dep in $(DEPS) ; do $(GO) get -u $@ ; done
