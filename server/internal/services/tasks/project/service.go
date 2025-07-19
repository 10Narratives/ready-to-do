package projectsrv

import (
	"context"
	"fmt"

	projectmodels "github.com/10Narratives/ready-to-do/server/internal/models/tasks/project"
	projectapi "github.com/10Narratives/ready-to-do/server/internal/transport/grpc/tasks/project"
	"google.golang.org/grpc/status"
)

//go:generate mockery --name ProjectStorage --output ./mocks/
type ProjectStorage interface {
	Create(ctx context.Context, project *projectmodels.Project) *status.Status
}

type Serice struct {
	storage ProjectStorage
}

var _ projectapi.ProjectService = &Serice{}

func ProjectName(projectID string) string {
	return fmt.Sprintf("project/%s", projectID)
}

func New(storage ProjectStorage) *Serice {
	return &Serice{
		storage: storage,
	}
}

func (s *Serice) Create(ctx context.Context, args projectapi.CreateProjectArgs) *status.Status {
	args.Project.Name = ProjectName(args.ProjectID)
	return s.storage.Create(ctx, args.Project)
}
