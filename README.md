# Setup Protobuf and gRPC
## Install protobuf compiler
For MacOS:
```bash
brew install protobuf
```
## Install protoc-gen-go
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Update your PATH so that the protoc compiler can find the plugins:
```bash
export PATH="$PATH:$(go env GOPATH)/bin" >> ~/.zshrc
```

# Setup gRPC Gateway
## Install protoc-gen-grpc-gateway
```bash
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```