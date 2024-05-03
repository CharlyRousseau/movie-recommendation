package database

import (
    "testing"
    "github.com/golang/mock/gomock"
    "movie-reccomendation-api/mocks"
    "movie-reccomendation-api/models"
)

func TestCreateUser(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabaseService(ctrl)
    user := &models.User{Username: "john", Email: "john@example.com", Password: "password"}

    mockDB.EXPECT().CreateUser(user).Return(nil).Times(1)

    if err := mockDB.CreateUser(user); err != nil {
        t.Errorf("Failed to create user: %v", err)
    }
}

func TestLoginUser(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabaseService(ctrl)
    userEmail := "john@example.com"
    userPassword := "password"
    expectedUser := &models.User{ID: 1, Username: "john", Email: userEmail, Password: userPassword}

    mockDB.EXPECT().LoginUser(userEmail, userPassword).Return(expectedUser, nil).Times(1)

    user, err := mockDB.LoginUser(userEmail, userPassword)
    if err != nil {
        t.Errorf("LoginUser failed: %v", err)
    }
    if user.Email != userEmail {
        t.Errorf("Expected user email %v, got %v", userEmail, user.Email)
    }
}

func TestGetUserByUsernameOrEmail(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabaseService(ctrl)
    username := "john"
    email := "john@example.com"
    expectedUser := &models.User{ID: 1, Username: username, Email: email}

    mockDB.EXPECT().GetUserByUsernameOrEmail(username, email).Return(expectedUser, nil).Times(1)

    user, err := mockDB.GetUserByUsernameOrEmail(username, email)
    if err != nil {
        t.Errorf("GetUserByUsernameOrEmail failed: %v", err)
    }
    if user.Username != username || user.Email != email {
        t.Errorf("Expected user %v/%v, got %v/%v", username, email, user.Username, user.Email)
    }
}

func TestGetUserByID(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabaseService(ctrl)
    userID := int64(1)
    expectedUser := &models.User{ID: userID, Username: "john", Email: "john@example.com"}

    mockDB.EXPECT().GetUserByID(userID).Return(expectedUser, nil).Times(1)

    user, err := mockDB.GetUserByID(userID)
    if err != nil {
        t.Errorf("GetUserByID failed: %v", err)
    }
    if user.ID != userID {
        t.Errorf("Expected user ID %v, got %v", userID, user.ID)
    }
}

func TestAddFavorite(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabaseService(ctrl)
    userID := int64(1)
    itemID := int64(100)
    
    mockDB.EXPECT().AddFavorite(userID, itemID).Return(nil).Times(1)

    if err := mockDB.AddFavorite(userID, itemID); err != nil {
        t.Errorf("AddFavorite failed: %v", err)
    }
}

func TestGetFavoritesByUser(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabaseService(ctrl)
    userID := int64(1)
    expectedFavorites := []models.Favorite{{ID: 1, UserID: userID, ItemID: 100}}

    mockDB.EXPECT().GetFavoritesByUser(userID).Return(expectedFavorites, nil).Times(1)

    favorites, err := mockDB.GetFavoritesByUser(userID)
    if err != nil {
        t.Errorf("GetFavoritesByUser failed: %v", err)
    }
    if len(favorites) != 1 || favorites[0].ItemID != 100 {
        t.Errorf("Expected 1 favorite with itemID 100, got %d favorites with itemID %v", len(favorites), favorites[0].ItemID)
    }
}

func TestGetFavoriteByUserAndId(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabaseService(ctrl)
    userID := int64(1)
    itemID := int64(100)
    expectedFavorite := &models.Favorite{ID: 1, UserID: userID, ItemID: itemID}

    mockDB.EXPECT().GetFavoriteByUserAndId(userID, itemID).Return(expectedFavorite, nil).Times(1)

    favorite, err := mockDB.GetFavoriteByUserAndId(userID, itemID)
    if err != nil {
        t.Errorf("GetFavoriteByUserAndId failed: %v", err)
    }
    if favorite.ItemID != itemID {
        t.Errorf("Expected favorite itemID %v, got %v", itemID, favorite.ItemID)
    }
}

func TestRemoveFavorite(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDB := mocks.NewMockDatabaseService(ctrl)
    userID := int64(1)
    itemID := int64(100)
    
    mockDB.EXPECT().RemoveFavorite(userID, itemID).Return(nil).Times(1)

    if err := mockDB.RemoveFavorite(userID, itemID); err != nil {
        t.Errorf("RemoveFavorite failed: %v", err)
    }
}
