package projectapi_test

import (
	"context"
	"testing"
	"time"

	tasksv1 "github.com/10Narratives/ready-to-do/contracts/gen/go/proto/tasks/v1"
	projectmodels "github.com/10Narratives/ready-to-do/server/internal/models/tasks/project"
	projectapi "github.com/10Narratives/ready-to-do/server/internal/transport/grpc/tasks/project"
	"github.com/10Narratives/ready-to-do/server/internal/transport/grpc/tasks/project/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerAPI_CreateProject(t *testing.T) {
	t.Parallel()

	var (
		projectID string = "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"
		project          = &projectmodels.Project{
			Name:        "projects/" + projectID,
			DisplayName: "the awesome project",
			Description: "the awesome decription",
			ColorTag:    "#000000",
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			State:       projectmodels.ActiveProjectState,
		}
		invalidProject = &projectmodels.Project{
			Name:        "invalid project name",
			DisplayName: "the awesome project",
			Description: "the awesome decription",
			ColorTag:    "#000000",
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			State:       1000,
		}
	)

	type fields struct {
		setupProjectServiceMock func(m *mocks.ProjectService)
	}
	type args struct {
		ctx context.Context
		req *tasksv1.CreateProjectRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    require.ValueAssertionFunc
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "successful execution",
			fields: fields{
				setupProjectServiceMock: func(m *mocks.ProjectService) {
					m.On("Create", mock.Anything, projectapi.CreateProjectArgs{
						ProjectID: projectID,
						Project:   project,
					}).Return(nil)
				},
			},
			args: args{
				ctx: context.Background(),
				req: &tasksv1.CreateProjectRequest{
					ProjectId: projectID,
					Project:   projectmodels.ProjectToGRPC(project),
				},
			},
			want: func(tt require.TestingT, got interface{}, _ ...interface{}) {
				p, ok := got.(*tasksv1.Project)
				require.True(t, ok)

				converted, err := projectmodels.ProjectFromGRPC(p)
				require.NoError(t, err)

				assert.Equal(t, project, converted)
			},
			wantErr: require.NoError,
		},
		{
			name: "validation error recieved",
			fields: fields{
				setupProjectServiceMock: func(m *mocks.ProjectService) {
				},
			},
			args: args{
				ctx: context.Background(),
				req: &tasksv1.CreateProjectRequest{
					ProjectId: "projectID",
					Project:   projectmodels.ProjectToGRPC(invalidProject),
				},
			},
			want:    require.Empty,
			wantErr: require.Error,
		},
		{
			name: "unsupported project state",
			fields: fields{
				setupProjectServiceMock: func(m *mocks.ProjectService) {
				},
			},
			args: args{
				ctx: context.Background(),
				req: &tasksv1.CreateProjectRequest{
					ProjectId: projectID,
					Project:   projectmodels.ProjectToGRPC(invalidProject),
				},
			},
			want:    require.Empty,
			wantErr: require.Error,
		},
		{
			name: "failed precondition",
			fields: fields{
				setupProjectServiceMock: func(m *mocks.ProjectService) {
					m.On("Create", mock.Anything, projectapi.CreateProjectArgs{
						ProjectID: projectID,
						Project:   project,
					}).Return(status.New(codes.FailedPrecondition, "some failed precondition"))
				},
			},
			args: args{
				ctx: context.Background(),
				req: &tasksv1.CreateProjectRequest{
					ProjectId: projectID,
					Project:   projectmodels.ProjectToGRPC(project),
				},
			},
			want:    require.Empty,
			wantErr: require.Error,
		},
		{
			name: "internal error",
			fields: fields{
				setupProjectServiceMock: func(m *mocks.ProjectService) {
					m.On("Create", mock.Anything, projectapi.CreateProjectArgs{
						ProjectID: projectID,
						Project:   project,
					}).Return(status.New(codes.Internal, "some intrenal error"))
				},
			},
			args: args{
				ctx: context.Background(),
				req: &tasksv1.CreateProjectRequest{
					ProjectId: projectID,
					Project:   projectmodels.ProjectToGRPC(project),
				},
			},
			want:    require.Empty,
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			projectServiceMock := mocks.NewProjectService(t)
			tt.fields.setupProjectServiceMock(projectServiceMock)

			api := projectapi.New(projectServiceMock)
			resp, err := api.CreateProject(tt.args.ctx, tt.args.req)

			tt.want(t, resp)
			tt.wantErr(t, err)
		})
	}
}
