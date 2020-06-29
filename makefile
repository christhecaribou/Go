 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

PROGRAM=./src

all: test build
build:
	$(GOBUILD) $(PROGRAM)
test: 
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)