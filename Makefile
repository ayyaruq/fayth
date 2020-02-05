# Go params
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOGEN=$(GOCMD) generate
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt

all: build
build:
	$(GOGEN) ./opcodes
	$(GOFMT) ./opcodes
clean:
	$(GOCLEAN)
deps:
	$(GOGET) -u

.PHONY: clean all
