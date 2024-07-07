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

	middleware.SetupLoggerMiddleware(e)

	db := config.InitDB()
	pc := &controllers.PortfolioController{DB: db}

	config.SetupTemplates(e)
	routes.SetupRoutes(e, pc)

	e.Logger.Fatal(e.Start(":8080"))
}
