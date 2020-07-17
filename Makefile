GOPATH=$(shell pwd)/vendor:$(shell pwd)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")
PID=/tmp/go-$(GONAME).pid
DESTDIR=$
PREFIX ?= /usr/local

.PHONY: build
build:
	@echo "Building $(GOFILES) to ./bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -ldflags="-X github.com/diegomagdaleno/cheemit/lib.prefix=${PREFIX}" -o bin/$(GONAME) $(GOFILES) 

.PHONY: get
get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

.PHONY: install
install:
	$(MAKE) -C .
	@install -d $(DESTDIR)$(PREFIX)/bin
	@install -m 777 ./bin/cheemit $(DESTDIR)$(PREFIX)/bin
	@install -d $(DESTDIR)$(PREFIX)/share/cheemit/font/
	@install -m 777 ./resources/fonts/Anton-Regular.ttf $(DESTDIR)$(PREFIX)/share/cheemit/font/
	@install -d $(DESTDIR)$(PREFIX)/share/cheemit/image
	@install -m 777 ./resources/images/Cheems.png $(DESTDIR)$(PREFIX)/share/cheemit/image
	@install -m 777 ./resources/images/Doge.png $(DESTDIR)$(PREFIX)/share/cheemit/image

.PHONY: run
run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

.PHONY: watch
watch:
	@$(MAKE) restart &
	@fswatch -o . -e 'bin/.*' | xargs -n1 -I{}  make restart

	restart: clear stop clean build start

.PHONY: start
start:
	@echo "Starting bin/$(GONAME)"
	@./bin/$(GONAME) & echo $$! > $(PID)

.PHONY: stop
stop:
	@echo "Stopping bin/$(GONAME) if it's running"
	@-kill `[[ -f $(PID) ]] && cat $(PID)` 2>/dev/null || true

.PHONY: clear
clear:
	@clear

.PHONY: clean
clean:
	@echo "Cleaning"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean -modcache

.PHONY: vendor
vendor:
	@go mod init \
			&& go mod tidy \
			&& go mod vendor
