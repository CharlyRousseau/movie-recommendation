package routes

import (
	"movie-reccomendation-api/database"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *database.DB,) {
	e.POST("/users", CreateUser(db))
	e.POST("/login", LoginUser(db))
	e.POST("/favorites", AddFavoriteHandler(db))
    e.GET("/favorites", GetFavoritesHandler(db))
}
