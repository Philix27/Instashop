include .env
export $(shell sed 's/=.*//' .env)

# Variables
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_RUN_DEV=$(GO_CMD) run
BINARY_NAME=mobarter

# Default target
all: build

# Build the executable
build:
	$(GO_BUILD) -o $(BINARY_NAME) -v

# Clean up generated files
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)

# Run the executable
run:
	$(GO_BUILD) -o $(BINARY_NAME) -v ./
	./$(BINARY_NAME)

# Run development
rd:
	swag fmt && swag init && $(GO_CMD) fmt && $(GO_RUN_DEV) ./main.go

# Install dependencies
deps:
	$(GO_GET) -v ./

# Run tests
test:
	$(GO_TEST) -v ./

# Run 'go fmt' on all go files
fmt:
	$(GO_CMD) fmt ./... && swag fmt

# Run 'go vet' to check for common mistakes
vet:
	$(GO_CMD) vet ./

# Run all checks
check: fmt vet test

db_gen: 
	sqlc generate

db_new: 
	migrate create -ext sql -dir db/migrations -seq $(name) 
# make db_new name=enums

db_push: 
	migrate -database ${DB_URL} -path db/migrations up
# make db_push    

db_down: 
	migrate -database ${DB_URL} -path db/migrations down
# make db_push    



# Shortcuts
.PHONY: build clean run deps test fmt vet check

