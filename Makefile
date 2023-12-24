.PHONY: help checkhealth bootstrap build-dev download test build run install uninstall clean

# Default target executed when no arguments are given to make.
help:
	@echo "Available commands:"
	@echo "  checkhealth  	- Verify development dependencies are installed"
	@echo "  bootstrap  	- Install developer dependencies"
	@echo "  build-dev  	- Build the developer container"
	@echo "  download   	- Download dependencies"
	@echo "  test       	- Run tests"
	@echo "  build      	- Build the application"
	@echo "  run        	- Run the application (requires args 'FILEPATH|DRAFTER' and 'SEASON')"
	@echo "  install    	- Install the binary to OPATH/bin"
	@echo "  uninstall  	- Uninstall the binary from OPATH/bin"
	@echo "  clean      	- Remove built application and any generated files"

# Verify development tools are installed
checkhealth:
	@./script/checkhealth.sh

# Install tools from tools.go
bootstrap:
	@cat tools.go | grep _ | awk '{ print $$2 }' | xargs -L1 go install
	@cp script/commit-msg.sh .git/hooks/commit-msg

# Build the developer container
build-dev:
	docker build -f Dockerfile.dev -t apple-health-report-dev .

# Download necessary dependencies
download:
	go mod download
	go mod tidy

# Run tests
test:
	gotest -v ./...

# Build the app
build:
	go build -o apple-health-report ./cmd

# Run the app
run:
	AHR_LOG_LEVEL=DEBUG ./apple-health-report

# Install the binary to $GOPATH/bin
install:
	go build -o $(GOPATH)/bin/apple-health-report

# Uninstall the binary from $GOPATH/bin
uninstall:
	rm -f $(GOPATH)/bin/apple-health-report

# Clean up
clean:
	rm -f apple-health-report

# Set the default goal to 'help' when no targets are given on the command line
.DEFAULT_GOAL := help

