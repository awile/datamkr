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
	
build_mac_amd:
	GOOS=darwin GOARCH=amd64 go build

build_mac_arm:
	GOOS=darwin GOARCH=arm64 go build
