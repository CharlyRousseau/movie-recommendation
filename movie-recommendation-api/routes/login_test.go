package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"movie-reccomendation-api/mocks"
	"movie-reccomendation-api/models"
)

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockDatabaseService(mockCtrl)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(`{"username": "testuser", "email": "test@example.com", "password": "testpassword"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	// Mock database to return nil, indicating that the user does not exist
	mockDB.EXPECT().GetUserByUsernameOrEmail("testuser", "test@example.com").Return(nil, nil)
	// Mock database to return nil, indicating no error in user creation
	mockDB.EXPECT().CreateUser(gomock.Any()).Return(nil)

	// Call CreateUser handler function
	if assert.NoError(t, CreateUser(mockDB)(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestLoginUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocks.NewMockDatabaseService(mockCtrl)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email": "test@example.com", "password": "testpassword"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")

	// Mock database to return a user with the provided credentials
	mockUser := &models.User{ID: 1, Username: "testuser", Email: "test@example.com"}
	mockDB.EXPECT().LoginUser("test@example.com", "testpassword").Return(mockUser, nil)

	// Call LoginUser handler function
	if assert.NoError(t, LoginUser(mockDB)(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Unmarshal the response JSON and assert the response
		var responseMap map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.NotNil(t, responseMap["user"])
		assert.NotNil(t, responseMap["token"])
	}
}
