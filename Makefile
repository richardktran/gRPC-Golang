gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=. --go-grpc_out=.
run-server:
	go run calculator/server/server.go
run-client:
	go run calculator/client/client.go