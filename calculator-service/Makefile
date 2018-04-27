build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/go-calculator/calculator-service \
	  proto/calculator/calculator.proto

	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t calculator-service	.

run:
	sudo docker run -p 9001:9001 calculator-service 
