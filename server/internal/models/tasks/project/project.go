package projectmodels

import (
	"errors"
	"fmt"
	"time"

	tasksv1 "github.com/10Narratives/ready-to-do/contracts/gen/go/proto/tasks/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProjectState string

const (
	UnspecifiedProjectState ProjectState = "unspecified"
	ActiveProjectState      ProjectState = "active"
	ArchivedProjectState    ProjectState = "archived"
	DeletedprojectState     ProjectState = "deleted"
)

func ProjectStateFromGRPC(src tasksv1.Project_State) (ProjectState, error) {
	switch src {
	case tasksv1.Project_ARCHIVED:
		return ArchivedProjectState, nil
	case tasksv1.Project_ACTIVE:
		return ActiveProjectState, nil
	case tasksv1.Project_DELETED:
		return DeletedprojectState, nil
	default:
		return UnspecifiedProjectState, errors.New("unsupported project state")
	}
}

func ProjectStateToGRPC(src ProjectState) tasksv1.Project_State {
	switch src {
	case ArchivedProjectState:
		return tasksv1.Project_ARCHIVED
	case ActiveProjectState:
		return tasksv1.Project_ACTIVE
	case DeletedprojectState:
		return tasksv1.Project_DELETED
	default:
		return tasksv1.Project_STATE_UNSPECIFIED
	}
}

type Project struct {
	Name        string       `json:"name"`
	DisplayName string       `json:"display_name"`
	Description string       `json:"description"`
	ColorTag    string       `json:"color_tag"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	State       ProjectState `json:"state"`
}

func ProjectFromGRPC(src *tasksv1.Project) (*Project, error) {
	if src == nil {
		return nil, nil
	}

	state, err := ProjectStateFromGRPC(src.GetState())
	if err != nil {
		return nil, fmt.Errorf("cannot convert project model from grpc: %v", err)
	}

	return &Project{
		Name:        src.GetName(),
		DisplayName: src.GetDisplayName(),
		Description: src.GetDescription(),
		ColorTag:    src.GetColorTag(),
		CreatedAt:   src.GetCreatedAt().AsTime(),
		UpdatedAt:   src.GetUpdatedAt().AsTime(),
		State:       state,
	}, nil
}

func ProjectToGRPC(src *Project) *tasksv1.Project {
	if src == nil {
		return nil
	}

	return &tasksv1.Project{
		Name:        src.Name,
		DisplayName: src.DisplayName,
		Description: src.Description,
		ColorTag:    src.ColorTag,
		CreatedAt:   timestamppb.New(src.CreatedAt),
		UpdatedAt:   timestamppb.New(src.UpdatedAt),
		State:       ProjectStateToGRPC(src.State),
	}
}
