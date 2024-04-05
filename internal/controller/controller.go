package controller

import (
	"fmt"
	"login_app/internal/models"
	"login_app/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	UserService  services.UserService
	LoginService services.LoginService
}

func (contr *Controller) CreateUser(c echo.Context) error {
	currUser := models.User{}
	err := c.Bind(&currUser)
	if err != nil {
		return err
	}

	err = contr.UserService.CreateUser(c.Request().Context(), currUser)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "")
}

func (contr *Controller) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := contr.UserService.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	return c.NoContent(http.StatusAccepted)
}

func (contr *Controller) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	token, err := contr.LoginService.Login(c.Request().Context(), username, password)
	fmt.Printf("token: %s\n", token)
	c.Response().Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (contr *Controller) Validate(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	err := contr.LoginService.Validate(c.Request().Context(), token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Valid")
}

func NewController(userService services.UserService, loginService services.LoginService) *Controller {
	return &Controller{
		userService,
		loginService,
	}
}
