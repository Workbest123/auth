package main

import (
    "github.com/Workbest123/auth/internal/config"
      "log/slog"
      "os"
      "github.com/Workbest123/auth/internal/app"
)

const (
    envLocal = "local"
    envDev   = "dev"
    envProd  = "prod"
)

func main(){
    cfg := config.MustLoad("./config/local.yaml")
    log := setupLogger(cfg.Env)
    log.Info("starting application", slog.Any("config",cfg))
    application:=app.New(log,cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
    application.GRPCSrv.Mustrun()
}


func setupLogger(env string) *slog.Logger {
    var log *slog.Logger

    switch env {
    case envLocal:
        log = slog.New(
            slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envDev:
        log = slog.New(
            slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envProd:
        log = slog.New(
            slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
        )
    }

    return log
}