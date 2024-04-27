package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"login_app/internal/config"
	"login_app/internal/controller"
	"login_app/internal/db"
	"login_app/internal/repository"
	"login_app/internal/services"
	"os"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrEpiredToken  = errors.New("expired token")
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

	r := e.Group("/api")
	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningMethod: jwt.SigningMethodRS256.Name,

		KeyFunc: func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodRS256.Name {
				return nil, ErrInvalidToken
			}

			key, err := os.ReadFile(config.EnvConfigs.PUBLIC_SIGN_KEY_PATH)
			if err != nil {
				return nil, err
			}
			pem, err := jwt.ParseRSAPublicKeyFromPEM(key)
			if err != nil {
				return nil, err
			}

			return pem, nil
		},
		//ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
		//	encryptedToken := c.Request().Header.Get("Authorization")
		//	clearToken := strings.Split(encryptedToken, "Bearer ")
		//	parse, err := jwt.ParseWithClaims(clearToken[0], models.JwtTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//		return util.GetPublicSignKey()
		//	})
		//	if err != nil {
		//		return nil, err
		//	}
		//	return parse, nil
		//},
	}))
	r.GET("/test", func(c echo.Context) error {
		return c.String(200, "Success")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
