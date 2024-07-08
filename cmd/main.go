package main

import (
	"bookswapper/internal/app"
	//	_ "bookswapper/cmd/docs"
)

// @title	Bookswapper API
// @version		1.0
// @description	API for offline book exchange service
// @host	localhost:8080
func main() {
	// create app
	bookSwapperApp := app.NewBookswapperApp()

	// start app
	app.Start(bookSwapperApp)
}
