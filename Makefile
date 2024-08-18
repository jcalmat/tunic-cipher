# Go compiler
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
APP_NAME := tunic
BIN_FOLDER := bin

# Default target
all: test build

# Build the binary
build:
	$(GOBUILD) -o $(APP_NAME) ./...

tidy:
	$(GOCMD) mod tidy

vendor:
	$(GOCMD) mod vendor

run:
	$(GOCMD) run ./...
# Run tests
test:
	$(GOTEST) -v ./...

# Clean up binaries
clean:
	$(GOCLEAN)
	rm -f $(APP_NAME) $(BIN_FOLDER)/$(APP_NAME)

# Install dependencies
deps:
	$(GOGET) -v ./...

# Cross compile for windows and linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) ./...

build-windows:
	env GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOARCH=amd64 $(GOBUILD) -o $(BIN_FOLDER) ./...
	cp -R resources $(BIN_FOLDER)/

.PHONY: all build clean test deps build-linux
