package workergrpc

import (
	"context"

	ddbv1 "github.com/10Narratives/contracts/gen/go/ddb"
	"github.com/10Narratives/ddb/internal/domain/models"
	"google.golang.org/grpc"
)

type Worker interface {
	CreateDocument(ctx context.Context, collectionID string, payload map[string]string) (models.Document, error)
	ReadDocument(ctx context.Context, collectionID, documentID string) (models.Document, error)
	ReadAllDocuments(ctx context.Context, collectionID string) ([]models.Document, error)
	UpdateDocument(ctx context.Context, collectionID, documentID string, fieldsToUpdate map[string]string) error
	DeleteDocument(ctx context.Context, collectionID, documentID string) error
}

type serverAPI struct {
	ddbv1.UnimplementedWorkerServer
	worker Worker
}

func Register(gRPC *grpc.Server, worker Worker) {
	ddbv1.RegisterWorkerServer(gRPC, &serverAPI{worker: worker})
}

func (s *serverAPI) Create(ctx context.Context, req *ddbv1.CreateRequest) (*ddbv1.CreateResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Read(ctx context.Context, req *ddbv1.ReadRequest) (*ddbv1.ReadResponse, error) {
	panic("implement me")
}

func (s *serverAPI) ReadAll(ctx context.Context, req *ddbv1.ReadAllRequest) (*ddbv1.ReadAllResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Update(ctx context.Context, req *ddbv1.UpdateRequest) (*ddbv1.UpdateResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Delete(ctx context.Context, req *ddbv1.DeleteRequest) (*ddbv1.DeleteResponse, error) {
	panic("implement me")
}
