GO=go

all: plugins go-plugin-example

go-plugin-example: main.go
	$(GO) build

plugins:
	$(MAKE) -C plugins

clean:
	$(MAKE) -C plugins clean
	$(GO) clean

.PHONY: plugins clean
