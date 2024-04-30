package routes

import (
	"net/http"

	"movie-reccomendation-api/database"

	"github.com/labstack/echo/v4"
)

func CreateUser(db *database.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user database.User
		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		existingUser, err := db.GetUserByUsernameOrEmail(user.Username, user.Email)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if existingUser != nil {
			return echo.NewHTTPError(http.StatusConflict, "Username or email is already taken")
		}

		if err := db.CreateUser(&user); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}
