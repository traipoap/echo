// main.go
package main

import (
	"echo/config"
	"echo/controllers"
	"echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := config.InitDB()
	pc := &controllers.PortfolioController{DB: db}

	config.SetupTemplates(e)
	routes.SetupRoutes(e, pc)

	e.Logger.Fatal(e.Start(":8080"))
}
