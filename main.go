// main.go
package main

import (
	"echo/config"
	"echo/controllers"
	"echo/middleware"
	"echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Serve static files
	e.Static("/assets", "views/assets")
	e.Static("/static", "views/static")
	e.Static("/.html", "views")

	middleware.SetupLoggerMiddleware(e)

	db := config.InitDB()
	pc := &controllers.PortfolioController{DB: db}

	config.SetupTemplates(e)
	routes.SetupRoutes(e, pc)

	e.Logger.Fatal(e.Start(":8080"))
}
