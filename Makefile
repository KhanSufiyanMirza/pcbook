gen:
	 protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
	#protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb --grpc-gateway_out=:pb --openapiv2_out=:swagger
.PHONY: gen