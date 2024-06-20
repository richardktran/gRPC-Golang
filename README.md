# Install
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
