build_linux:
	@GOOS=linux GOARCH=amd64 go build -v -o bin/send-linux ./examples/send

build: 
	@go build -o bin/send ./examples/send