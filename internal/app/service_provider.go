package app

import (
	"github.com/Bdido86/auth/internal/api/user"
)

type serviceProvider struct {
	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) UserImpl() *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation()
	}

	return s.userImpl
}
