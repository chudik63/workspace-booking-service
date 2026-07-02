package workspace

import (
	"context"

	model "github.com/golangmonster/workspace-booking-service/internal/model/workspace"
	dto "github.com/golangmonster/workspace-booking-service/internal/service/workspace"
)

type workspaceService interface {
	GetWorkspaceByID(ctx context.Context, id int64) (*model.Workspace, error)
	CreateWorkspace(ctx context.Context, req *dto.CreateWorkspaceRequest) (*model.Workspace, error)
	UpdateWorkspace(ctx context.Context, req *dto.UpdateWorkspaceRequest) error
	DeleteWorkspaceByID(ctx context.Context, id int64) error
	ListWorkspaces(ctx context.Context, req *dto.ListWorkspacesRequest) (*dto.ListWorkspacesResponse, error)
}
