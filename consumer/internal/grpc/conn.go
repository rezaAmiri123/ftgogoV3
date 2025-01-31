package grpc

import (
	"context"

	"github.com/stackus/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func clientErrorUnrayInterceptor()grpc.UnaryClientInterceptor{
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return errors.ReceiveGRPCError(invoker(ctx,method, req,reply,cc,opts...))
	}
}

func Dial(ctx context.Context, endpoint string) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(clientErrorUnrayInterceptor()),		
	)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			if err = conn.Close(); err != nil {
				// TODO do something when logging is a thing
			}
		}
		go func() {
			<-ctx.Done()
			if err = conn.Close(); err != nil {
				// TODO do something when logging is a thing
			}
		}()
	}()
	return
}
