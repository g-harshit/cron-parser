# Sample Makefile

# Define the default target
all: build

# Rule for building the project
build:
	go build -o myapp main.go

# Rule for running tests
test:
	go test ./...
