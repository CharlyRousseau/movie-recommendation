package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"movie-reccomendation-api/models"

)

type DatabaseInterface interface {
    CreateUser(user *models.User) error
    LoginUser(email, password string) (*models.User, error)
    GetUserByUsernameOrEmail(username, email string) (*models.User, error)
    GetUserByID(id int64) (*models.User, error)
    AddFavorite(userID, itemID int64) error
    GetFavoritesByUser(userID int64) ([]models.Favorite, error)
    GetFavoriteByUserAndId(userID, itemID int64) (*models.Favorite, error)
    RemoveFavorite(userID, itemID int64) error
}

type DB struct {
	*gorm.DB
}

func NewDB(connectionString string) (DatabaseInterface, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DB) LoginUser(email, password string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid login credentials")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid login credentials")
	}

	return &user, nil
}

func (db *DB) GetUserByUsernameOrEmail(username, email string) (*models.User, error) {
	var user models.User
	err := db.Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (db *DB) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func (db *DB) AddFavorite(userID, itemID int64) error {
	favorite := models.Favorite{
		UserID: userID,
		ItemID: itemID,
	}

	result := db.Create(&favorite)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetFavoritesByUser(userID int64) ([]models.Favorite, error) {
	var favorites []models.Favorite
	result := db.Where("user_id = ?", userID).Find(&favorites)
	if result.Error != nil {
		return nil, result.Error
	}

	return favorites, nil
}
func (db *DB) GetFavoriteByUserAndId(userID, itemID int64) (*models.Favorite, error) {
	var favorite models.Favorite
	result := db.Where("user_id = ? AND item_id = ?", userID, itemID).First(&favorite)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &favorite, nil
}

func (db *DB) RemoveFavorite(userID, itemID int64) error {
	result := db.Where("user_id = ? AND item_id = ?", userID, itemID).Delete(&models.Favorite{})
	return result.Error
}
