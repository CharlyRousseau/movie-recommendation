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
	ID     int64
	UserID int64
	ItemID int64
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

func (db *DB) GetUserByID(id int64) (*User, error) {
	var user User
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
	favorite := Favorite{
		UserID: userID,
		ItemID: itemID,
	}

	result := db.Create(&favorite)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *DB) GetFavoritesByUser(userID int64) ([]Favorite, error) {
	var favorites []Favorite
	result := db.Where("user_id = ?", userID).Find(&favorites)
	if result.Error != nil {
		return nil, result.Error
	}

	return favorites, nil
}
func (db *DB) GetFavoriteByUserAndId(userID, itemID int64) (*Favorite, error) {
	var favorite Favorite
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
	result := db.Where("user_id = ? AND item_id = ?", userID, itemID).Delete(&Favorite{})
	return result.Error
}
