install:
	go install github.com/awile/datamkr

fmt:
	gofmt -s -w .

lint:
	golangci-lint run
