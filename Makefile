GO=go
TARGET=plugintest

all: plugins $(TARGET)

$(TARGET): plugintest.go
	$(GO) build -o $@

plugins:
	$(MAKE) -C plugins

clean:
	$(MAKE) -C plugins clean
	$(GO) clean

.PHONY: plugins clean
