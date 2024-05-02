package routes

import (
	"movie-reccomendation-api/database"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Id int64 `json:"name"`
	jwt.RegisteredClaims
}

func RegisterRoutes(e *echo.Echo, db *database.DB) {
	e.POST("/users", CreateUser(db))
	e.POST("/login", LoginUser(db))
	r := e.Group("")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/me", GetMeHandler(db))
	r.GET("/favorites", GetFavorites(db))
	r.POST("/favorites", AddFavorite(db))
	r.POST("/favorites/remove", RemoveFavorite(db))
}
