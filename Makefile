BINARY_NAME = swipe
BIN_DIR = bin
CMD_DIR = cmd
MAIN_FILE = $(CMD_DIR)/main.go

TARGET_DIR = $(HOME)/.local/bin
INSTALL_DIR = $(TARGET_DIR)/$(BINARY_NAME)

build:
	@echo "Building the project..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_FILE)

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

# Build and run the project
all: build install run