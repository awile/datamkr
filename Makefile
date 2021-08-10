install:
	go install github.com/awile/datamkr
	
check:
	make fmt && make lint && make test 

test:
	go test ./...

fmt:
	gofmt -s -w .

lint:
	golangci-lint run
