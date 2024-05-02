package routes

import (
	"movie-reccomendation-api/database"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("your_secret_key") // Use a more secure key in production

func LoginUser(db *database.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(database.User)

		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := db.LoginUser(user.Email, user.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid login credentials")
		}

		// Create token
		token, err := generateJWT(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"user":  user,
			"token": token,
		})
	}
}

func generateJWT(user *database.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   strconv.FormatInt(user.ID, 10),
		ExpiresAt: expirationTime.Unix(),
		Issuer:    "movie-recommendation",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
