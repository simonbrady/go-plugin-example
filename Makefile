GO=go

all: plugins go-plugin-example

go-plugin-example:
	$(GO) build

plugins:
	$(MAKE) -C plugins

clean:
	$(MAKE) -C plugins clean
	$(GO) clean

.PHONY: go-plugin-example plugins clean
