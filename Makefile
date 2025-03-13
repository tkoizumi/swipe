BINARY_NAME = swipe
BIN_DIR = bin
CMD_DIR = cmd
MAIN_FILE = $(CMD_DIR)/main.go
UTILS_FILE = $(CMD_DIR)/utils.go

TARGET_DIR = $(HOME)/.local/bin
INSTALL_DIR = $(TARGET_DIR)/$(BINARY_NAME)
GO_VERSION := $(shell go version | cut -d ' ' -f 3 | sed 's/go//')
GO_REQUIRED_VERSION := 1.22.3
ARCH := $(shell uname -m)

# Build and run the project
all: ensure_go build install

ensure_go:
	@echo "Checking go version..."
	@echo "CPU Architecture: $(ARCH)"
	@echo "Current Go version: $(GO_VERSION)"
	@if [ -n "$(GO_VERSION)" ]; then \
			if [ "$(GO_VERSION)" = "$(GO_REQUIRED_VERSION)" ]; then \
					echo "Go version is 1.22.3"; \
			else \
					echo "Go version is not 1.22.3. Please ensure Go version is exactly 1.22.3"; \
					curl -sSL https://go.dev/dl/go${GO_REQUIRED_VERSION}.darwin-amd64.tar.gz | tar -xz -C $(HOME)/.go --strip-components=1; \
			fi; \
	else \
			echo "Go is not installed or not in PATH. Please install Go."; \
	fi

build:
	@echo "Building the project..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_FILE) $(UTILS_FILE)

run: build
	@echo "Running the project..."
	./$(BIN_DIR)/$(BINARY_NAME)


install: build
	@echo "Installing swipe..."
	mkdir -p $(TARGET_DIR)
	@cp $(BIN_DIR)/$(BINARY_NAME) $(INSTALL_DIR)
	
	@echo "Installation complete"

clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)