# Linux focused Makefile

clean:
	@echo "Killing any process on port 8080"
	-fuser -k 8080/tcp || true

run: clean
	@echo "Running the application"
	export GIN_MODE=release && go run main.go

run-80:
	@echo "Running the application on port 80"
	export GIN_MODE=release && export PORT=80 && sudo -E go run main.go

build:
	go build -o bin/main main.go

format:
	go fmt ./...

test:
	go test -v ./...