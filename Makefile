up:
	docker build --progress=plain -t submit-api .
	docker run submit-api

run:
	go run .

unit-test:
	go test ./... -v -short

test:
	go test ./... -v

code-coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out | grep total | awk '{print $3}'

lint:
	golangci-lint run
