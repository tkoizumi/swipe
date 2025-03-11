BINARY_NAME = swipe
BIN_DIR = bin
CMD_DIR = cmd
MAIN_FILE = $(CMD_DIR)/main.go
UTILS_FILE = $(CMD_DIR)/utils.go

TARGET_DIR = $(HOME)/.local/bin
INSTALL_DIR = $(TARGET_DIR)/$(BINARY_NAME)

GO_VERSION = 1.22.3
ARCH := $(shell uname -m)
GO_URL := https://go.dev/dl/go$(GO_VERSION).darwin-$(if $(filter arm64,$(ARCH)),arm64,amd64).tar.gz
GO_DIR = $(HOME)/.go
GO_CMD = $(GO_DIR)/bin/go

# Build and run the project
all: ensure_go build install

ensure_go:
	@echo "Checking Go version..."
	@if [ ! -x "$(GO_CMD)" ] || ! $(GO_CMD) version 2>/dev/null | grep -q "go$(GO_VERSION)"; then \
		echo "Installing Go $(GO_VERSION) for $(ARCH)..."; \
		rm -rf $(GO_DIR); \
		mkdir -p $(GO_DIR); \
		curl -sSL $(GO_URL) | tar -xz -C $(GO_DIR) --strip-components=1; \
	fi
	@echo "Go $(GO_VERSION) is installed at $(GO_CMD)."

build: ensure_go
	@echo "Building the project..."
	mkdir -p $(BIN_DIR)
	$(GO_CMD) build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_FILE) $(UTILS_FILE)

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
	@echo "Cleaning up...
	rm -rf $(BIN_DIR)

