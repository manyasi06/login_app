package controller

import (
	"context"
	"errors"
	"login_app/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockUserService struct{}
type mockLoginService struct{}

func (m *mockLoginService) Login(ctx context.Context, username, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

// create mockLoginService implements LoginService interface methods
func (m *mockLoginService) Validate(ctx context.Context, token string) error {
	return nil
}

func (m *mockUserService) DeleteUser(ctx context.Context, id string) error {
	if id == "existing_id" {
		return nil
	}
	return errors.New("user not found")
}

func (m *mockUserService) CreateUser(ctx context.Context, user models.User) error {
	return nil
}

func TestController_DeleteUserBadRequest(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("non_existing_id")

	userService := &mockUserService{}
	loginService := &mockLoginService{} // You can create a mock login service if needed
	controller := NewController(userService, loginService)

	err := controller.DeleteUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "{\"message\":\"user not found\"}\n", rec.Body.String())
}

func TestController_DeleteUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("existing_id")

	userService := &mockUserService{}
	loginService := &mockLoginService{} // You can create a mock login service if needed
	controller := NewController(userService, loginService)

	// Test case: Delete existing user
	err := controller.DeleteUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusAccepted, rec.Code)
	assert.Equal(t, "", rec.Body.String())
}
