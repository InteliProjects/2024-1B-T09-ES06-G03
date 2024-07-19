build:
	@go build -o bin/api cmd/main.go

# test:
# 	@go test -v ./...

# less verbose
test:
	@for dir in $(shell find ./service -type d); do \
		if [ -n "$$(ls $$dir/*.go 2>/dev/null)" ]; then \
			echo "Running tests in $$dir"; \
			go test $$dir; \
		fi \
	done

run: build
	@./bin/api


