package main

// This is just a placeholder application at this time.  I will implement something more robust, soon.

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"

	// Use the following imports when generating from proto files using 'buf generate'.
	//helloworldv1 "github.com/abitofhelp/bzlmod/internal/gen/abitofhelp/helloworld/v1"
	//"github.com/abitofhelp/bzlmod/internal/gen/abitofhelp/helloworld/v1/helloworldv1connect"

	// Use the following imports when generating from proto files using Bazel build.
	helloworldv1 "github.com/abitofhelp/bzlmod/bazel-bin/proto/abitofhelp/helloworld/v1/abitofhelp_helloworld_v1_go_proto_/github.com/abitofhelp/bzlmod/proto/abitofhelp/helloworld/v1"
	// FIXME: "github.com/abitofhelp/bzlmod/bazel-bin/proto/abitofhelp/helloworld/v1/abitofhelp_helloworld_v1_go_proto_/github.com/abitofhelp/bzlmod/proto/abitofhelp/helloworld/v1/helloworldconnect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// NewGreeterServer returns a new GreeterServer implementation
func NewGreeterServer() helloworldv1connect.GreeterServiceHandler {
	return &server{}
}

// grpcurl -plaintext localhost:8080 list
// grpcurl -plaintext -d '{"name": "Jane"}' localhost:8080 abitofhelp.helloworld.v1.GreeterService/SayHello
// curl -X POST -H "Content-Type: application/json" -d '{"name": "Jane"}' "http://127.0.0.1:8080/abitofhelp.helloworld.v1.GreeterService/SayHello"
// buf curl --http2-prior-knowledge  -d '{"name": "Jane"}' "http://127.0.0.1:8080/abitofhelp.helloworld.v1.GreeterService/SayHello" -v

func (x *server) SayHello(ctx context.Context, req *connect.Request[helloworldv1.SayHelloRequest]) (*connect.Response[helloworldv1.SayHelloResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&helloworldv1.SayHelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greeter-Version", "v1")
	return res, nil
}

func main() {
	s := NewGreeterServer()
	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector(
		helloworldv1connect.GreeterServiceName,
	)

	// Many tools still expect the older version of the server reflection API, so
	// most servers should mount both handlers.
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	mux.Handle(grpcreflect.NewHandlerV1(reflector))

	path, handler := helloworldv1connect.NewGreeterServiceHandler(s)
	mux.Handle(path, handler)
	err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		return
	}
}
