package routes

import (
	"movie-reccomendation-api/database"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// GetFavorites returns a handler function that retrieves all favorites for a user
func GetFavorites(db database.DatabaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwtCustomClaims)
		userId := claims.Id

		favorites, err := db.GetFavoritesByUser(userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, favorites)
	}
}

func AddFavorite(db database.DatabaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwtCustomClaims)
		userId := claims.Id

		itemId, err := strconv.ParseInt(c.QueryParam("itemId"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid itemId")
		}

		existingFavorite, err := db.GetFavoriteByUserAndId(userId, itemId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if existingFavorite != nil {
			return echo.NewHTTPError(http.StatusConflict, "Favorite already exists")
		}
		err = db.AddFavorite(userId, itemId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, map[string]string{"message": "Favorite added successfully"})
	}
}
func RemoveFavorite(db database.DatabaseInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwtCustomClaims)
		userId := claims.Id

		itemId, err := strconv.ParseInt(c.QueryParam("itemId"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid itemId")
		}

		err = db.RemoveFavorite(userId, itemId)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Favorite removed successfully"})
	}
}
