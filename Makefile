APP_NAME=dealbreakerapi
BUILD_DIR=./bin
ENTRYPOINT=./api/main.go

.PHONY: all clean build run

all: clean build run

clean: 
	@rm -rf $(BUILD_DIR)
	
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(ENTRYPOINT)

run: 
	@$(BUILD_DIR)/$(APP_NAME)	
