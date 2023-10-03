package user

import (
	"context"
	"log/slog"

	"github.com/Bdido86/auth/pkg/api/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/brianvoe/gofakeit"
)

func (i *Implementation) Get(_ context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	slog.Info("method Create input params:", "id", req.GetId())

	return &user_v1.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      0,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}
