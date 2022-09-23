PROJECT_NAME=restful-api-demo
MAIN_FILE=main.go

dep: ## Get the dependencies
	@go mod tidy

build: dep ## Build the binary file
	@go build -ldflags "-s -w" -o dist/api-demo $(MAIN_FILE)

run: dep ## Run develop server
	@go run $(MAIN_FILE) start -f etc/demo.toml

linux: dep ## Build the binary file for linux
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/api-demo $(MAIN_FILE)

clean: ## Remove all binary file
	@rm -f dist/*