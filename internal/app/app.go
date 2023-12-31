package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/Bdido86/auth/pkg/api/user_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// todo add config
const url = "localhost:5051"

type App struct {
	grpcServer      *grpc.Server
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	app := new(App)

	err := app.init(ctx)
	if err != nil {
		return nil, fmt.Errorf("newApp init error: %w", err)
	}

	return app, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) init(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return fmt.Errorf("newApp init error: %w", err)
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	user_v1.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserImpl())

	return nil
}

func (a *App) runGRPCServer() error {
	slog.Info("GRPC server is starting on:", "url", url)

	list, err := net.Listen("tcp", url)
	if err != nil {
		return fmt.Errorf("runGRPCServer Listen tcp error: %w", err)
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return fmt.Errorf("runGRPCServer Serve error: %w", err)
	}

	return nil
}
