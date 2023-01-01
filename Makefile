build:
	mkdir -p functions
	go build -o ./functions/question ./cmd/question/main.go
	go build -o ./functions/answer   ./cmd/answer/main.go