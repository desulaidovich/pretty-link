package main

import "github.com/desulaidovich/pretty-link/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
