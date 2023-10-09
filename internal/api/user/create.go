package user

import (
	"context"
	"log/slog"

	"github.com/Bdido86/auth/pkg/api/user_v1"
)

func (i *Implementation) Create(_ context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	var result int64

	slog.Info("method Create input params:", "name", req.GetName(), "email", req.GetEmail())

	return &user_v1.CreateResponse{
		Id: result,
	}, nil
}
