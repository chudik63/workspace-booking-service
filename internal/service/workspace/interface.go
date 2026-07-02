package workspace

import (
	"context"

	model "github.com/golangmonster/workspace-booking-service/internal/model/workspace"
)

type workspaceRepository interface {
	GetWorkspaceByID(ctx context.Context, id int64) (*model.Workspace, error)
	InsertWorkspace(ctx context.Context, req *CreateWorkspaceRequest) (*model.Workspace, error)
	UpdateWorkspace(ctx context.Context, req *UpdateWorkspaceRequest) error
	DeleteWorkspaceByID(ctx context.Context, id int64) error
	ListWorkspaces(ctx context.Context, req *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	InTx(ctx context.Context, f func(ctx context.Context) error) error
}
