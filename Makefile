# Directories
CMD_DIR = ./cmd/gig-agg
BIN_DIR = bin

# Target executable
TARGET = $(BIN_DIR)/gig-agg.out

# Default target
all: build

# Build the target executable
build:
	@mkdir -p $(BIN_DIR)
	@cp -r ./assets $(BIN_DIR) 
	go build -o $(TARGET) $(CMD_DIR)

# Run the application
run: build
	$(TARGET)

# Clean build files
clean:
	rm -rf $(BIN_DIR)

# Phony targets
.PHONY: all build run clean
