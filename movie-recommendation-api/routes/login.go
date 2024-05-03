package routes

import (
	"movie-reccomendation-api/database"
	"movie-reccomendation-api/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func CreateUser(db database.DatabaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
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

func LoginUser(db database.DatabaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := db.LoginUser(user.Email, user.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		claims := &jwtCustomClaims{
			user.ID,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		userResponse := &UserResponse{
			ID:    user.ID,
			Email: user.Email,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"user":  userResponse,
			"token": t,
		})
	}
}
