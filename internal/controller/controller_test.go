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

func (m *mockLoginService) Login(username, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

// create mockLoginService implements LoginService interface methods
func (m *mockLoginService) Validate(token string) error {
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

func TestController_DeleteUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/existing_id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userService := &mockUserService{}
	loginService := &mockLoginService{} // You can create a mock login service if needed
	controller := NewController(userService, loginService)

	// Test case: Delete existing user
	err := controller.DeleteUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusAccepted, rec.Code)
	assert.Equal(t, "", rec.Body.String())

	// Test case: Delete non-existing user
	req = httptest.NewRequest(http.MethodDelete, "/users/non_existing_id", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err = controller.DeleteUser(c)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "user not found\n", rec.Body.String())
}
