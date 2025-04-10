# Directories
GIG_AGG_CMD_DIR = ./cmd/gig-agg
TUI_CMD_DIR = ./cmd/tui
BIN_DIR = bin

# Target executable
GIG_AGG_TARGET = $(BIN_DIR)/gig-agg.out
TUI_TARGET = $(BIN_DIR)/tui.out

# Default target
all: build

# Build the target executable
build:
	@mkdir -p $(BIN_DIR)
	@cp -r ./assets $(BIN_DIR) 
	go build -o $(GIG_AGG_TARGET) $(GIG_AGG_CMD_DIR)
	go build -o $(TUI_TARGET) $(TUI_CMD_DIR)
	@echo "Build complete. Executables are located in $(BIN_DIR)."

# Run the application
run_server: build
	$(GIG_AGG_TARGET)

run_tui: build
	$(TUI_TARGET)

# Clean build files
clean:
	rm -rf $(BIN_DIR)

# Phony targets
.PHONY: all build run clean
