package main

import (
	"github.com/labstack/echo/v4"
	"login_app/internal/config"
	"login_app/internal/controller"
	"login_app/internal/db"
	"login_app/internal/repository"
	"login_app/internal/services"
)

func main() {
	//Init configs
	config.InitEnvConfigs()
	db.InitDB()

	userRepository := repository.NewUserRepository(db.DBClient.Client.Database("user"))
	loginRepository := repository.NewLoginRepository(db.DBClient.Client.Database("user"))

	userServ := services.NewUserService(userRepository)
	loginServ := services.NewLoginService(loginRepository)

	controller := controller.NewController(userServ, loginServ)

	e := echo.New()
	e.POST("/user", controller.CreateUser)
	e.DELETE("/user/:id", controller.DeleteUser)
	e.POST("/login", controller.Login)
	e.GET("/validate", controller.Validate)

	e.Logger.Fatal(e.Start(":1323"))

}
