package workspace

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/golangmonster/workspace-booking-service/pkg/api/workspace/v1"
)

type Implementation struct {
	pb.UnimplementedWorkspaceServiceServer

	workspaceService workspaceService
}

func New(service workspaceService) *Implementation {
	return &Implementation{
		workspaceService: service,
	}
}

func (i *Implementation) RegisterServer(server *grpc.Server) {
	pb.RegisterWorkspaceServiceServer(server, i)
}

func (i *Implementation) RegisterHandlerFromEndpoint(
	ctx context.Context,
	mux *runtime.ServeMux,
	addrGRPC string,
	opts []grpc.DialOption,
) error {
	return pb.RegisterWorkspaceServiceHandlerFromEndpoint(ctx, mux, addrGRPC, opts)
}
