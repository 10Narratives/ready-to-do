package projectstore

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/status"

	projectmodels "github.com/10Narratives/ready-to-do/server/internal/models/tasks/project"
	projectsrv "github.com/10Narratives/ready-to-do/server/internal/services/tasks/project"
)

type Storage struct {
}

var _ projectsrv.ProjectStorage = &Storage{}

func New(db *sql.DB) *Storage {
	return &Storage{}
}

func (s *Storage) Create(ctx context.Context, project *projectmodels.Project) *status.Status {
	return nil
}
