format:
	@echo "Check files for formatting"
	@gofmt -l -w .

check:
	@echo "Checking source code for suspicious activity"
	@go vet .

build:
	@echo "Building executable"
	@go build -o mml ./main.go