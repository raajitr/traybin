build:
	@go build -o bin/traybin

run: build
	@./bin/traybin

test:
	@go test -v ./..