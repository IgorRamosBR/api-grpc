# API gRPC Web

## Generating gRPC code

### Installing protocol buffer compiler

`go get -u github.com/golang/protobuf/protoc-gen-go`

`go get -u google.golang.org/grpc`

### Generating code

`protoc --go_out=plugins=grpc:. ./pkg/proto/users.proto`