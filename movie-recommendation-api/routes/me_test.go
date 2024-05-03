package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/golang-jwt/jwt/v4"

	"movie-reccomendation-api/mocks"
	"movie-reccomendation-api/models"
)

func TestGetMeHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockDatabaseService(mockCtrl)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/me")

	// Mock JWT claims
	claims := &jwtCustomClaims{Id: 1}
	token := &jwt.Token{Claims: claims}
	c.Set("user", token)

	// Mock database to return a user with the provided ID
	mockUser := &models.User{ID: 1, Username: "testuser", Email: "test@example.com"}
	mockDB.EXPECT().GetUserByID(claims.Id).Return(mockUser, nil)

	// Call GetMeHandler handler function
	if assert.NoError(t, GetMeHandler(mockDB)(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Unmarshal the response JSON and assert the response
		var responseUser models.User
		err := json.Unmarshal(rec.Body.Bytes(), &responseUser)
		if assert.NoError(t, err) {
			assert.Equal(t, mockUser, &responseUser)
		}
	}
}
