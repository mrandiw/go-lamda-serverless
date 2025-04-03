# Run locally
run:
	@echo "Running locally..."
	IS_LOCAL=true go run main.go

build:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./bootstrap .
	zip main.zip bootstrap