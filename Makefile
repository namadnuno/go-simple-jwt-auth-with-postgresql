dev:
	go run main.go
db:
	go run main.go -seed -migrate
build:
	go build
