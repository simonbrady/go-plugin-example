GO=go
TARGET=plugintest

all: plugins $(TARGET)

$(TARGET): plugintest.go
	$(GO) build -o $@

plugins: hello
	$(MAKE) -C plugins

hello:
	$(MAKE) -C hello

clean:
	$(MAKE) -C hello clean
	$(MAKE) -C plugins clean
	$(GO) clean

.PHONY: plugins hello clean
