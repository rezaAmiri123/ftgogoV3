package rpc

import (
	"context"

	"github.com/stackus/errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func clientErrorUnaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return errors.ReceiveGRPCError(invoker(ctx, method, req, reply, cc, opts...))
	}
}

func Dial(endpoint string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			otelgrpc.UnaryClientInterceptor(),
			clientErrorUnaryInterceptor(),
		),
	)
	return grpc.NewClient(endpoint,opts...)
}
