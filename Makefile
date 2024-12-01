gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/ticketservice.proto
	
start_server:
	go run cmd/server/main.go

start_client:
	go run cmd/client/main.go