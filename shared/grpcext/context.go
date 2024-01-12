package grpcext

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"mabna/shared/contextext"
)

func LoadContext() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for key, value := range md {
				ctx = contextext.SetValue(ctx, key, value[0])
			}
		}
		return handler(ctx, req)
	}
}
