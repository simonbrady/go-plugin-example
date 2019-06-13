GO=go
BUILD=$(GO) build

TARGET=plugintest

all: plugins $(TARGET)

$(TARGET): plugintest.go
	$(BUILD) -o $@

plugins:
	$(MAKE) -C plugins

clean:
	$(MAKE) -C plugins clean
	-rm $(TARGET)

.PHONY: plugins clean
