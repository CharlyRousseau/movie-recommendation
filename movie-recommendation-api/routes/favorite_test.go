package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4/middleware"
	"github.com/golang-jwt/jwt/v4"

	"movie-reccomendation-api/models"
	"movie-reccomendation-api/mocks"
)

type JWTClaims struct {
    Id int64 `json:"name"`
    jwt.RegisteredClaims
}

type JWTToken struct {
    Claims *JWTClaims
}

func TestGetFavorites(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    mockDB := mocks.NewMockDatabaseService(mockCtrl)
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/favorites", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetPath("/favorites")

    // Mock JWT claims
    userClaims := &jwtCustomClaims{Id: 1}
    token := &jwt.Token{Claims: userClaims}
    c.Set("user", token)

    expectedFavorites := []models.Favorite{
        {ID: 1, UserID: 1, ItemID: 100},
    }

    mockDB.EXPECT().GetFavoritesByUser(userClaims.Id).Return(expectedFavorites, nil)

    // Call GetFavorites handler function
    if assert.NoError(t, GetFavorites(mockDB)(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        // You can also check the response body if necessary
        assert.Equal(t, expectedFavorites, rec.Body)
    }
}

func TestAddFavorite(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockDatabaseService(mockCtrl)
	e := echo.New()
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}))

	itemID := 100
	req := httptest.NewRequest(http.MethodPost, "/favorites", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer dummy")
	q := req.URL.Query()
	q.Add("itemId", strconv.Itoa(itemID))
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/favorites")

	// Mock JWT claims
	userClaims := JWTClaims{Id: 1}
	token := JWTToken{Claims: &userClaims}
	c.Set("user", token)

	mockDB.EXPECT().GetFavoriteByUserAndId(userClaims.Id, int64(itemID)).Return(nil, nil)
	mockDB.EXPECT().AddFavorite(userClaims.Id, int64(itemID)).Return(nil)

	if assert.NoError(t, AddFavorite(mockDB)(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		responseMap := make(map[string]string)
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Equal(t, "Favorite added successfully", responseMap["message"])
	}
}

func TestRemoveFavorite(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockDatabaseService(mockCtrl)
	e := echo.New()
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}))

	itemID := 100
	req := httptest.NewRequest(http.MethodPost, "/favorites/remove", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer dummy")
	q := req.URL.Query()
	q.Add("itemId", strconv.Itoa(itemID))
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/favorites/remove")

	// Mock JWT claims
	userClaims := JWTClaims{Id: 1}
	token := JWTToken{Claims: &userClaims}
	c.Set("user", token)

	mockDB.EXPECT().RemoveFavorite(userClaims.Id, int64(itemID)).Return(nil)

	if assert.NoError(t, RemoveFavorite(mockDB)(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		responseMap := make(map[string]string)
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.Equal(t, "Favorite removed successfully", responseMap["message"])
	}
}
