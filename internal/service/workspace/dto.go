package workspace

import (
	"github.com/golangmonster/workspace-booking-service/internal/model/page"
	model "github.com/golangmonster/workspace-booking-service/internal/model/workspace"
)

type CreateWorkspaceRequest struct {
	Name        string
	Lat         float64
	Lon         float64
	FullAddress string
	Type        model.Type
	Status      model.Status
	Capacity    uint32
}

type UpdateWorkspaceRequest struct {
	ID          int64
	Name        string
	Lat         float64
	Lon         float64
	FullAddress string
	Capacity    uint32
	Status      model.Status
}

type ListWorkspacesRequest struct {
	Page   *page.Page
	Filter *WorkspaceFilter
}

type WorkspaceFilter struct {
	Types       []model.Type
	Statuses    []model.Status
	CapacityGte *uint32
	FullAddress *string
}

type ListWorkspacesResponse struct {
	Workspaces []*model.Workspace
	TotalCount uint32
}
