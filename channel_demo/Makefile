GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=godemo

all: test run

run: 
	$(GORUN) .

test: 
	$(GOTEST) -v ./... -timeout 5000ms

build: 
	mkdir -p .bin
	$(GOBUILD) -o .bin/$(BINARY_NAME) -v

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)