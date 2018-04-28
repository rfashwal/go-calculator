build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/go-calculator/calculator-service \
	  proto/calculator/calculator.proto

	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t calculator-service -f Dockerfile.server	.

run:
sudo docker run -p 50051:50051 calculator-service 
