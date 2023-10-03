package user

import (
	desc "github.com/Bdido86/auth/pkg/api/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
}

func NewImplementation() *Implementation {
	return &Implementation{}
}
