build:
	mkdir -p functions
	go build -o ./functions/question ./cmd/question/main.go