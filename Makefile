GO=go

all: plugins go-plugin-example

go-plugin-example:
	$(GO) build

lint:
	staticcheck -checks all ./...

plugins:
	$(MAKE) -C plugins

clean:
	$(MAKE) -C plugins clean
	$(GO) clean

.PHONY: go-plugin-example lint plugins clean
