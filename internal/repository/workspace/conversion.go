package workspace

import model "github.com/golangmonster/workspace-booking-service/internal/model/workspace"

func toWorkspace(i workspaceItem) *model.Workspace {
	return &model.Workspace{
		ID:          i.ID,
		Name:        i.Name,
		Lat:         i.Lat,
		Lon:         i.Lon,
		FullAddress: i.FullAddress,
		Type:        model.Type(i.Type),
		Status:      model.Status(i.Status),
		Capacity:    i.Capacity,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		IsDeleted:   i.IsDeleted,
	}
}
