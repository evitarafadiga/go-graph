# Variables
GO = go
PROJECT_DIR = $(shell pwd)
GQLGEN = github.com/99designs/gqlgen

# Env

PORT = 8080

build:
	@echo "==>🔨 (1/2) Building server..."
	@$(GO) build -o $(PROJECT_DIR)/bin/server $(PROJECT_DIR)/server.go
	@echo "==> (2/2) Success! Built into: $(PROJECT_DIR)/bin/server"
run: build
	@echo "==>🏃‍♂️‍➡️ (1/1) Running server..."
	@PORT=$(PORT) $(PROJECT_DIR)/bin/server
generate:
	@echo "==> (1/2) Generating GraphQL Schema..."
	@$(GO) run $(GQLGEN) generate
	@echo "==> (2/2) Success! Schemas generated successfully!"
clean: 
	@echo "==> (1/2) Cleaning..."
	@rm -rf $(PROJECT_DIR)/bin
	@echo "==> (2/2) Cleaned."

.DEFAULT_GOAL := run