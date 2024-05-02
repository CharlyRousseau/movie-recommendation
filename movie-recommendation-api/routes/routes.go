package routes

import (
	"movie-reccomendation-api/database"
    "movie-reccomendation-api/tmdb" 

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *database.DB, tmdbClient *tmdb.Client) {
	e.POST("/users", CreateUser(db))
	e.POST("/login", LoginUser(db))
	e.GET("/movies", GetTrendingMoviesHandler(tmdbClient))
}
