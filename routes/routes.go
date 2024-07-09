// routes/routes.go
package routes

import (
	"echo/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, pc *controllers.PortfolioController) {
	e.GET("/", pc.Home)

	e.GET("/works", pc.Works)
	e.GET("/project-1", pc.Project1)
	e.GET("/project-2", pc.Project2)

	e.GET("/experience", pc.Experience)

	e.GET("/blog", pc.Blog)

	e.GET("/about", pc.About)

	e.GET("/stats", pc.GetVisitStats)

	// Middleware for recording visits
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			pc.RecordVisit(c)
			return next(c)
		}
	})
}
