package routes

import (
	"net/http"

	"movie-reccomendation-api/database"

	"github.com/labstack/echo/v4"
)

func LoginUser(db *database.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(database.User)

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := db.LoginUser(user.Email, user.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	}
}
