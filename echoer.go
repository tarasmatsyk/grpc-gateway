package echoer

import (
	"context"
	"flag"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"

	"gateway/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:9090", "gRPC server endpoint")
)

type srv struct {
	proto.EchoerServer
}

func (s srv) Echo(ctx context.Context, message *proto.StringMessage) (*proto.StringMessage, error) {
	return message, nil
}

func RunGRPC(ctx context.Context, listen string) error {
	ln, err := net.Listen("tcp", listen)
	if err != nil {
		return err
	}
	g := grpc.NewServer()
	proto.RegisterEchoerServer(g, &srv{})
	return g.Serve(ln)
}

func RunHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var g errgroup.Group
	g.Go(func() error {
		return RunGRPC(ctx, *grpcServerEndpoint)
	})

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	err := proto.RegisterEchoerHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	g.Go(func() error {
		return http.ListenAndServe(":8081", mux)
	})

	return g.Wait()
}
