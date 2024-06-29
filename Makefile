# Linux focused Makefile

clean:
	@echo "Killing any process on port 8080"
	-fuser -k 8080/tcp || true

run: clean
	@echo "Running the application"
	export GIN_MODE=release && go run main.go

build:
	go build -o bin/main main.go

format:
	go fmt ./...

test:
	go test -v ./...