build: 
	go build -o build/app main.go

test:
	go test -v ./...

run: 
	go run main.go
