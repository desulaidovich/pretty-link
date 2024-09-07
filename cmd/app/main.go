package main

import (
	"log/slog"
	"os"

	"github.com/desulaidovich/pretty-link/app"
	"github.com/desulaidovich/pretty-link/config"
)

func main() {
	options := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, options))

	if err := run(logger); err != nil {
		logger.Error("Ooops",
			"Error", err.Error(),
		)
	}
}

func run(_ *slog.Logger) error {
	cfg, err := config.LoadFromEnv()

	if err != nil {
		return err
	}

	application := app.New(
		app.WithConfig(cfg),
	)

	if err := application.Serve(); err != nil {
		return err
	}

	return nil
}
