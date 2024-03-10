# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o ./tmp/main ./cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Generate templates
templ:
	@echo "Generating templates..."
	@templ generate --watch

# Generate Tailwind CS
tw:
	@echo "Watching tailwind classes..."
	@tailwindcss -i tailwind.base.css -o web/static/styles.css --watch

# Test the application
test:
	@echo "Testing..."
	@APP_ENV=test go test tests/*_test.go -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
