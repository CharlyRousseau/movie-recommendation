package interfaces

import "movie-reccomendation-api/models"

type DatabaseService interface {
    CreateUser(user *models.User) error
    LoginUser(email, password string) (*models.User, error)
    GetUserByUsernameOrEmail(username, email string) (*models.User, error)
    GetUserByID(id int64) (*models.User, error)
    AddFavorite(userID, itemID int64) error
    GetFavoritesByUser(userID int64) ([]models.Favorite, error)
    GetFavoriteByUserAndId(userID, itemID int64) (*models.Favorite, error)
    RemoveFavorite(userID, itemID int64) error
}
