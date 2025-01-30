# Variables
GO = go
PROJECT_DIR = $(shell pwd)
GQLGEN = github.com/99designs/gqlgen
DB_URL = localhost:5432

# Env

PORT = 8080

goinstall:
	@echo "==>⬇️ (1/2) Downloading GO..."
	@yum update
	@yum upgrade
	@wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz
	@tar -xvf go1.23.5.linux-amd64.tar.gz
	@rm -rf /usr/local/go
	@mv go /usr/local
	@export GOROOT=/usr/local/go
	@export GOPATH=$$HOME/Apps/go-graph
	@export PATH=$$GOPATH/bin:$$GOROOT/bin:$$PATH
	@echo "==>⬇️ (2/2) GO installed."
build: goinstall
	@echo "==>🔨 (1/2) Building server..."
	@$(GO) build -o $(PROJECT_DIR)/bin/server $(PROJECT_DIR)/server.go
	@echo "==> (2/2) Success! Built into: $(PROJECT_DIR)/bin/server"
run: build
	@echo "==>🏃‍♂️‍➡️ (1/1) Running server..."
	@DB_NAME=$(DB_NAME) DB_URL=$(DB_URL) PORT=$(PORT) $(PROJECT_DIR)/bin/server
generate:
	@echo "==> (1/2) Generating GraphQL Schema..."
	@$(GO) run $(GQLGEN) generate
	@echo "==> (2/2) Success! Schemas generated successfully!"
clean: 
	@echo "==> (1/2) Cleaning..."
	@rm -rf $(PROJECT_DIR)/bin
	@echo "==> (2/2) Cleaned."

.DEFAULT_GOAL := run