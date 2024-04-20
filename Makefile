# Make sure the Makefile uses tabs for indentation
.PHONY: all go-build go-run vue-build vue-serve clean

# Default command when you run just `make`
all: go-build vue-build

# Build the Go project
go-build:
	@echo "Building Go project..."
	@go build -o bin/server main.go
# Run the Go server
go-run:
	@echo "Running Go server..."
	@./bin/server

# Build the Vue.js project
vue-build:
	@echo "Building Vue.js project..."
	@cd frontend && npm run build

# Serve the Vue.js project using the development server
vue-serve:
	@echo "Starting Vue.js development server..."
	@cd frontend && npm run serve

# Clean up generated files
clean:
	@echo "Cleaning up..."
	@rm -rf .crypto-tracker-frontend/dist
	@rm -rf bin

