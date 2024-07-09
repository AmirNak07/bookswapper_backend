package main

import (
	"bookswapper/internal/app"
)

func main() {
	// create app
	bookSwapperApp := app.NewBookswapperApp()

	// start app
	app.Start(bookSwapperApp)
}
