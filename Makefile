gen:
	 protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
	#protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb --openapiv2_out=:swagger
clean:
	rm pb/*.go
runServer:
	go run cmd/server/main.go -port 13400
runClient:
	go run cmd/client/main.go -address 0.0.0.0:13400
test:
	go test -cover -race ./...
	
.PHONY: gen,clean,run,test