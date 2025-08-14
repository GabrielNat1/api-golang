.PHONY: build run test clean

# Build variables
BINARY_NAME=api-server
BUILD_DIR=build

# Enable CGO for SQLite
export CGO_ENABLED=1

build:
	@echo "Building with CGO enabled..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./src/cmd/main.go

run:
	@CGO_ENABLED=1 go run ./src/cmd/main.go

test:
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean

dev:
	@CGO_ENABLED=1 go run ./src/cmd/main.go

install:
	@go mod download

# Cross compilation targets
build-windows:
	@echo "Building for Windows..."
	@set CGO_ENABLED=1
	@set GOOS=windows
	@set GOARCH=amd64
	@go build -o $(BUILD_DIR)/$(BINARY_NAME).exe ./src/cmd/main.go

build-linux:
	@echo "Building for Linux..."
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
	go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux ./src/cmd/main.go
