build:
	go build -o bin/bookstore main.go
run: build
	./bin/bookstore
dev:
	go run main.go
install:
	go install