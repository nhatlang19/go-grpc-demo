proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculatorpb/calculator.proto
start-server:
	go run server/server.go
start-client:
	go run client/client.go
start-client2:
	go run client/client-2.go