run:
	source .env && go run main.go

build:
	go build -o bin/gateway

test:
	go test ./...

totallines:
	find . -name '*.go' | xargs wc -l