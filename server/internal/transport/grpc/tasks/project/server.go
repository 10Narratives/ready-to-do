package projectapi

import (
	"context"
	"fmt"

	tasksv1 "github.com/10Narratives/ready-to-do/contracts/gen/go/proto/tasks/v1"
	projectmodels "github.com/10Narratives/ready-to-do/server/internal/models/tasks/project"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//go:generate mockery --name ProjectService --output ./mocks/
type ProjectService interface {
	// Create creates a new project using the provided arguments.
	// Returns the created project model and a gRPC status indicating the result.
	Create(ctx context.Context, args CreateProjectArgs) *status.Status
}

type CreateProjectArgs struct {
	ProjectID string
	Project   *projectmodels.Project
}

func newCreateProjectArgs(req *tasksv1.CreateProjectRequest) (CreateProjectArgs, error) {
	model, err := projectmodels.ProjectFromGRPC(req.GetProject())
	if err != nil {
		return CreateProjectArgs{}, fmt.Errorf("cannot convert request to create project args: %v", err)
	}
	return CreateProjectArgs{
		ProjectID: req.GetProjectId(),
		Project:   model,
	}, nil
}

type ServerAPI struct {
	tasksv1.UnimplementedProjectServiceServer
	service ProjectService
}

func New(service ProjectService) *ServerAPI {
	return &ServerAPI{
		service: service,
	}
}

func Register(server *grpc.Server, service ProjectService) {
	tasksv1.RegisterProjectServiceServer(server, &ServerAPI{service: service})
}

func (s *ServerAPI) CreateProject(ctx context.Context, req *tasksv1.CreateProjectRequest) (*tasksv1.Project, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	args, err := newCreateProjectArgs(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	if stat := s.service.Create(ctx, args); stat.Code() == codes.FailedPrecondition {
		return nil, status.Errorf(codes.FailedPrecondition, "failed precondition for project creation: %v", err)
	} else if stat != nil {
		return nil, status.Errorf(codes.Internal, "cannot create new project: %v", err)
	}

	return projectmodels.ProjectToGRPC(args.Project), nil
}

func (s *ServerAPI) DeleteProject(context.Context, *tasksv1.DeleteProjectRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *ServerAPI) GetProject(context.Context, *tasksv1.GetProjectRequest) (*tasksv1.Project, error) {
	return nil, nil
}

func (s *ServerAPI) ListProjects(context.Context, *tasksv1.ListProjectsRequest) (*tasksv1.ListProjectsResponse, error) {
	return nil, nil
}

func (s *ServerAPI) UpdateProject(context.Context, *tasksv1.UpdateProjectRequest) (*tasksv1.Project, error) {
	return nil, nil
}
