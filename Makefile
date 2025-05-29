APP_NAME     := myapp
BUILD_DIR    := bin
SRC          := ./src/cmd/main.go

GOCMD        := go
GOBUILD      := $(GOCMD) build
GOCLEAN      := $(GOCMD) clean
GOTEST       := $(GOCMD) test

.PHONY: all build run test clean

all: clean build run

build:
	@echo "→ Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(SRC)

run: build
	@echo "→ Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "→ Running tests..."
	$(GOTEST) ./...

clean:
	@echo "→ Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)