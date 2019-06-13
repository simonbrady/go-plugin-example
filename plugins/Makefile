GO=go

all: foo.so bar.so

clean:
	$(GO) clean

%.so: %.go
	$(GO) build -buildmode=plugin $<

.PHONY: clean
