package workspace

import (
	"context"
	"fmt"
)

func (s *service) UpdateWorkspace(ctx context.Context, req *UpdateWorkspaceRequest) error {
	err := s.repo.UpdateWorkspace(ctx, req)
	if err != nil {
		return fmt.Errorf("update workspace: %w", err)
	}

	return nil
}
