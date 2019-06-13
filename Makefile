GO=go

all: plugins go-plugin-example

go-plugin-example: main.go
	$(GO) build

plugins: hello
	$(MAKE) -C plugins

hello:
	$(MAKE) -C hello

clean:
	$(MAKE) -C hello clean
	$(MAKE) -C plugins clean
	$(GO) clean

.PHONY: plugins hello clean
