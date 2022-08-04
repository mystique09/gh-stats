package api

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mystique09/gh-profile/ent"
)

func Launch() {
	godotenv.Load()

	client, err := ent.Open("postgres", os.Getenv("DB_DB"))
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
	app.GET("/profile/:name", server.profileVisits)
	log.Fatal(app.Start(":3000"))
}
