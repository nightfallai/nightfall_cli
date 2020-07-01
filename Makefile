# Let Go know that our modules are private
export GOPRIVATE=github.com/watchtowerai

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

SERVICE_NAME=nightfall_dlp
BINARY_NAME=./$(SERVICE_NAME)
GO_TEST_ENV?=test

all: clean build start
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	GO_ENV=$(GO_TEST_ENV) $(GOTEST) ./... -count=1 -coverprofile=coverage.out
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
start:
	./$(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	go mod download
generate:
	go generate ./...
