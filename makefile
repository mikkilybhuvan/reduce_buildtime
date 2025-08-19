APP=bin/myapp

.PHONY: deps build run test clean
deps:
	go mod download

build:
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o $(APP) ./cmd/myapp

run: build
	./$(APP)

test:
	go test ./...

clean:
	rm -rf bin .gocache .gomodcache