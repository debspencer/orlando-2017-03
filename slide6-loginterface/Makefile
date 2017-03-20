SHELL    = /bin/bash
GOPATH   = $(shell pwd)
PLATFORM = $(shell uname)

GO        = GOPATH=$(GOPATH) /usr/local/go/bin/go

DEPS      = github.com/valyala/fasthttp \
            github.com/wadey/gocovmerge

ifeq ($(PLATFORM),Darwin)
BROWSER = open
else ifeq ($(PLATFORM),Linux)
BROWSER = x-www-browser
endif

DEPDIR  = $(patsubst %, src/%, $(DEPS))

SRCDIRS	= $(shell find src/goget -name '*.go' -print | while read file ; do echo `dirname $$file` ; done | sort -u)
DIRS = $(patsubst src/%, %, $(SRCDIRS))

all: clean $(DEPDIR) $(DIRS) goget.go
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
	@-rm -f cover.html cover.out
	@find . -name "cover.out" -delete

$(DIRS):
	@$(GO) vet $@
	@$(GO) test -tags $(shell basename $@) -coverprofile src/$@/cover.out $@


$(DEPDIR):
	$(GO) get -v $(patsubst src/%, %, $@)

update:
	for dep in $(DEPS) ; do $(GO) get -u $@ ; done

show: gocovmerge
	./bin/gocovmerge $(shell find src/goget -name cover.out -print) > cover.out
	$(GO) tool cover -html=cover.out -o cover.html
	$(BROWSER) cover.html
