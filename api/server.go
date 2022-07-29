package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Launch() {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	app.GET("/profile/:name", profileVisits)
	log.Fatal(app.Start(":3000"))
}
