package routes

import (
	"log"
	"net/http"
	"fmt"
	"movie-reccomendation-api/tmdb"
	"github.com/labstack/echo/v4"
)

// GetTrendingMoviesHandler récupère les films populaires depuis l'API TMDB et les renvoie au client.
func GetTrendingMoviesHandler(client *tmdb.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Effectue la requête vers l'API TMDB pour récupérer les films populaires
		movies, err := client.RequestAuthentication()
		fmt.Println(movies)
		if err != nil {
			log.Println("Failed to fetch trending movies:", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch trending movies")
		}
		
		// Renvoie les films populaires au client
		return c.JSON(http.StatusOK, movies)
	}
}

