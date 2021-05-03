GO=go
PLUGINS=foo bar

all: $(PLUGINS)

$(PLUGINS):
	$(GO) build -buildmode=plugin ./$@

clean:
	$(GO) clean

.PHONY: $(PLUGINS) clean
