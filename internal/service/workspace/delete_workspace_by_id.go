package workspace

import (
	"context"
	"fmt"
)

func (s *service) DeleteWorkspaceByID(ctx context.Context, id int64) error {
	err := s.repo.DeleteWorkspaceByID(ctx, id)
	if err != nil {
		return fmt.Errorf("delete workspace: %w", err)
	}

	return nil
}
