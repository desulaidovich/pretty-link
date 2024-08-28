package app

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/desulaidovich/pretty-link/auth/api"
	"github.com/desulaidovich/pretty-link/auth/usecase"
	"github.com/desulaidovich/pretty-link/config"
	"github.com/desulaidovich/pretty-link/middleware"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Run() error {
	logger := sloggerInit()

	cfg, err := config.LoadFromEnv()
	if err != nil {
		logger.Error("Failed from load config",
			slog.String("Message", err.Error()),
		)
		return err
	}

	db, err := databaseInit(cfg)

	if err != nil {
		logger.Error("Failed from database connection",
			slog.String("Message", err.Error()),
		)
		return err
	}

	defer db.Close()

	mux := http.NewServeMux()

	authUseCase := usecase.NewAuthUseCase(db)(logger)
	api.RegisterAuthEndpoints(mux, authUseCase)

	server := new(http.Server)
	server.Addr = ":" + cfg.Port
	server.Handler = middleware.New(logger)(mux)

	logger.Info("Appication "+cfg.Name+" starting",
		slog.String("port", cfg.Port),
	)

	if err = server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func databaseInit(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.ConnectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func sloggerInit() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)

	return slog.New(handler)
}
