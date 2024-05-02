package routes

import (
    "net/http"
    "strconv"
    "movie-reccomendation-api/database"
    "github.com/labstack/echo/v4"
)

func AddFavoriteHandler(db *database.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userID, err := strconv.ParseInt(c.FormValue("user_id"), 10, 64)
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
        }
        movieID, err := strconv.ParseInt(c.FormValue("movie_id"), 10, 64)
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid movie ID")
        }

        if err := db.AddFavorite(userID, movieID); err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }
        return c.JSON(http.StatusOK, "Favorite added successfully")
    }
}

func GetFavoritesHandler(db *database.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        userID, err := strconv.ParseInt(c.QueryParam("user_id"), 10, 64)
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
        }

        movies, err := db.GetFavorites(userID)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
        }
        return c.JSON(http.StatusOK, movies)
    }
}
