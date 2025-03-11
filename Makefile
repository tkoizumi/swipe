BINARY_NAME = swipe
BIN_DIR = bin
CMD_DIR = cmd
MAIN_FILE = $(CMD_DIR)/main.go
UTILS_FILE = $(CMD_DIR)/utils.go

TARGET_DIR = $(HOME)/.local/bin
INSTALL_DIR = $(TARGET_DIR)/$(BINARY_NAME)

GO_VERSION = 1.22.3
GO_BIN = $(HOME)/.go/bin/go$(GO_VERSION)
GO_CMD = $(HOME)/.go/bin/go

# Build and run the project
all: ensure_go build install

ensure_go:
	@echo "Checking Go version..."
	@if ! $(GO_CMD) version 2>/dev/null | grep -q 'go$(GO_VERSION)'; then \
		echo "Installing Go $(GO_VERSION)..."; \
		mkdir -p $(HOME)/.go && curl -sSL https://go.dev/dl/go$(GO_VERSION).darwin-amd64.tar.gz | tar -xz -C $(HOME)/.go --strip-components=1; \
	fi
	@echo "Go $(GO_VERSION) is installed."

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

