package main

import (
    "github.com/Workbest123/auth/internal/config"
      "log/slog"
      "os"
)

const (
    envLocal = "local"
    envDev   = "dev"
    envProd  = "prod"
)

func main(){
    cfg := config.MustLoad("./config/local.yaml")
    log := setupLogger(cfg.Env)
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