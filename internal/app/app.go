package app

import (
	"bookswapper/internal/api/routes"
	dbmodels "bookswapper/internal/models/database"
	"bookswapper/internal/utils/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// _ "bookswapper/cmd/docs"
)

type BookswapperApp struct {
	Fiber *fiber.App
	Db    *gorm.DB
}

func NewBookswapperApp() *BookswapperApp {
	// init db connection
	db, dbErr := database.Connection()
	if dbErr != nil {
		panic("failed to connect database" + dbErr.Error())
	}

	// migrate db models
	migrateErr := db.AutoMigrate(dbmodels.User{}, dbmodels.Book{}, dbmodels.Trade{})
	if migrateErr != nil {
		panic("failed to migrate database" + migrateErr.Error())
	}

	// init new fiber app and use swagger
	app := fiber.New()
	//app.Use(swagger.New())

	// map api routes and swagger
	//app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.Group("/api")
	routes.AuthRouter(api, db)
	routes.PingRouter(api)
	routes.ProfileRouter(api, db)

	return &BookswapperApp{
		Fiber: app,
		Db:    db,
	}
}

func Start(app *BookswapperApp) {
	if err := app.Fiber.Listen(":8080"); err != nil {
		panic("failed to listen: " + err.Error())
	}
}
