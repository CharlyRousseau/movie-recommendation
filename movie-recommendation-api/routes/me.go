package routes

import (
	"movie-reccomendation-api/database"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetMeHandler(db database.DatabaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		userId := claims.Id

		userDb, err := db.GetUserByID(userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user")
		}

		return c.JSON(http.StatusOK, userDb)
	}
}
