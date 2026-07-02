package workspace

import (
	"context"
	"fmt"
)

func (s *service) ListWorkspaces(ctx context.Context, req *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	resp, err := s.repo.ListWorkspaces(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("list workspaces: %w", err)
	}

	return resp, nil
}
