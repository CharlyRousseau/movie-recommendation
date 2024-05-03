package main

import (
	"log"
	"movie-reccomendation-api/config"
	"movie-reccomendation-api/database"
	"movie-reccomendation-api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.NewDB(cfg.DBConnectionString())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.CorsFrontEnd},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	routes.RegisterRoutes(e, db)

	log.Fatal(e.Start(":8080"))
}
