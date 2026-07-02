package workspace

type service struct {
	repo workspaceRepository
}

func New(repo workspaceRepository) *service {
	return &service{
		repo: repo,
	}
}
