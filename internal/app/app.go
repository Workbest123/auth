package app

import (
    "log/slog"
	"time"
	grpcapp "github.com/Workbest123/auth/internal/grpc"
)



	type App struct{
		GRPCServer *grpcapp.App
	}

	func New (
		log *slog.Logger,
		grpcPort int,
		storagePath string,
		tokenTTL time.Duration,
	)*App{
		grpcApp:= grpcapp.NewServer()
		return &App{
			GRPCServer: grpcApp,
		}
	}