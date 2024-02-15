build: 
	go build -o build/app main.go

test:
	go test -v ./...

lint:
	golangci-lint run

run: 
	go run main.go
