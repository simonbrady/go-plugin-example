GO=go
BUILD=$(GO) build

TARGETS=foo.so bar.so

all: $(TARGETS)

clean:
	-rm $(TARGETS)

%.so: %.go
	$(BUILD) -buildmode=plugin $^

.PHONY: clean
