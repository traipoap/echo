package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupLoggerMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, src=${remote_ip}, method=${method}, uri=${uri}, status=${status}, error=${error}, latency=${latency_human}\n",
	}))
}
