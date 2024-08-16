format:
	echo "Formatting files"
	gofmt -l -w .

check:
	echo "Checking source code for suspicious activity"
	go vet .