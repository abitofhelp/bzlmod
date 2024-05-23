module github.com/abitofhelp/bzlmod

go 1.22

// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

require (
	connectrpc.com/connect v1.16.2
	connectrpc.com/grpchealth v1.3.0
	connectrpc.com/grpcreflect v1.2.0
	golang.org/x/net v0.25.0
	google.golang.org/protobuf v1.34.1
)

require (
	golang.org/x/text v0.15.0 // indirect
	//google.golang.org/protobuf/cmd/protoc-gen-go v1.3.0 // indirect
    //google.golang.org/grpc/cmd/protoc-gen-go-grpc v0.0.0 // indirect
)
