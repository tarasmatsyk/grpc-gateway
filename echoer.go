package echoer

import (
	"context"
	"flag"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"net"
	"net/http"

	pr "gateway/proto/echoer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:9090", "gRPC server endpoint")
)

type srv struct {
	pr.EchoerServer
}

func (s srv) Echo(ctx context.Context, message *pr.StringMessage) (*pr.StringMessage, error) {
	return message, nil
}

func RunGRPC(ctx context.Context, listen string) error {
	ln, err := net.Listen("tcp", listen)
	if err != nil {
		return err
	}
	g := grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if rand.Intn(10) > 5 {
			return UnauthorizedInterceptor(ctx, req, info, handler)
		}

		return handler(ctx, req)
	}))
	pr.RegisterEchoerServer(g, &srv{})
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
	err := pr.RegisterEchoerHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	g.Go(func() error {
		return http.ListenAndServe(":8081", mux)
	})

	return g.Wait()
}
