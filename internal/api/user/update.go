package user

import (
	"context"
	"log/slog"

	"github.com/Bdido86/auth/pkg/api/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(_ context.Context, req *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	slog.Info("method Update input params:", "id", req.GetId(), "name", req.GetName(), "email", req.GetEmail())

	return &emptypb.Empty{}, nil
}
