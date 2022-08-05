package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mystique09/gh-profile/ent"
	"github.com/mystique09/gh-profile/web"
)

func Launch() {
	godotenv.Load()
	MODE := os.Getenv("MODE")
	PORT := os.Getenv("PORT")
	PORT = ":" + PORT

	if PORT == "" {
		PORT = ":3000"
	}

	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := Server{
		db: client,
	}
	app := echo.New()

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	app.Use(middleware.Recover())

	enableDevelopment(app, MODE)
	app.GET("/profile/:name", server.profileVisits, githubMiddleware)
	log.Fatal(app.Start(PORT))
}

func enableDevelopment(e *echo.Echo, mode string) {
	if mode == "development" {
		e.StaticFS("", web.BuildHTTPS())
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root: "web/dist",
			//Browse: true,
			HTML5: true,
		}))
	} else if mode == "production" {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Filesystem: http.FS(web.BuildHTTPS()),
			HTML5:      true,
		}))

	} else {
		log.Fatal(errors.New("invalid mode"))
	}
}
