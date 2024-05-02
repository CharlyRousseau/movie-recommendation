package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
}

type Favorite struct {
	UserID  int64 `gorm:"primaryKey;autoIncrement:false"`
	MovieID int64 `gorm:"primaryKey;autoIncrement:false"`
}

func NewDB(connectionString string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) CreateUser(user *User) error {
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

func (db *DB) LoginUser(email, password string) (*User, error) {
	var user User
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

func (db *DB) GetUserByUsernameOrEmail(username, email string) (*User, error) {
	var user User
	err := db.Where("username = ? OR email = ?", username, email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (db *DB) AddFavorite(userID, movieID int64) error {
	fav := Favorite{UserID: userID, MovieID: movieID}
	result := db.Create(&fav)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DB) GetFavorites(userID int64) ([]int64, error) {
	var favorites []Favorite
	result := db.Where("user_id = ?", userID).Find(&favorites)
	if result.Error != nil {
		return nil, result.Error
	}

	movieIDs := make([]int64, len(favorites))
	for i, fav := range favorites {
		movieIDs[i] = fav.MovieID
	}
	return movieIDs, nil
}
