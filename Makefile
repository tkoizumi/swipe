BINARY_NAME = swipe
BIN_DIR = bin
CMD_DIR = cmd
MAIN_FILE = $(CMD_DIR)/main.go
UTILS_FILE = $(CMD_DIR)/utils.go

TARGET_DIR = $(HOME)/.local/bin
INSTALL_DIR = $(TARGET_DIR)/$(BINARY_NAME)
GO_VERSION := $(shell go version | cut -d ' ' -f 3 | sed 's/go//')
GO_MAJOR_MINOR := $(shell echo $(GO_VERSION) | cut -d. -f1,2)
GO_REQUIRED_VERSION := 1.22
ARCH := $(shell uname -m)

# Build and run the project
all: ensure_go build install

ensure_go:
	@echo "Checking go version..."
	@echo "CPU Architecture: $(ARCH)"
	@echo "Current Go version: $(GO_VERSION)"
	@if [ -n "$(GO_VERSION)" ]; then \
			if awk 'BEGIN {exit !($(GO_MAJOR_MINOR) >= $(GO_REQUIRED_VERSION))}'; then \
					echo "Go version is compatible"; \
			else \
					echo "Go version must be at least 1.22. Please upgrade Go."; exit 1;\
			fi; \
	else \
			echo "Go is not installed or not in PATH. Please install Go."; exit 1;\
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