.PHONY: default run build test docs clean
# Variables
APP_NAME=acim-backend
MAIN_FILE=main.go

# Tasks
default: run-with-docs

run:
	@air
run-with-docs:
	@swag init
	@air
build:
	@go build -o $(APP_NAME) $(MAIN_FILE)
test:
	@go test ./...
docs:
	@swag init
clean:
	@rm -f (APP_NAME)
	@rm -f ./docs
setup:
	@go mod download
	@go get github.com/cosmtrek/air