package user

import (
	"context"
	"log/slog"

	"github.com/Bdido86/auth/pkg/api/user_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(_ context.Context, req *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	slog.Info("method Delete input params", "id", req.GetId())

	return &emptypb.Empty{}, nil
}
