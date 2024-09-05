package main

import (
	application "github.com/desulaidovich/pretty-link/app"
	"github.com/desulaidovich/pretty-link/config"
)

func main() {
	cfg, err := config.LoadFromEnv()

	if err != nil {
		panic(err)
	}

	app := application.New(
		application.WithConfig(cfg),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
