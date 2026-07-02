package workspace

import (
	"context"
	"fmt"

	model "github.com/golangmonster/workspace-booking-service/internal/model/workspace"
)

func (s *service) GetWorkspaceByID(ctx context.Context, id int64) (*model.Workspace, error) {
	workspace, err := s.repo.GetWorkspaceByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get workspace by id: %w", err)
	}

	return workspace, nil
}
