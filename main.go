package main

import (
	"fmt"
	"log"
	"todoApp/configs"
	_userHandler "todoApp/delivery/handler/user"
	_middleware "todoApp/delivery/middlewares"
	_routes "todoApp/delivery/routes"
	_userRepository "todoApp/repository/user"
	_userUseCase "todoApp/usecase/user"
	"todoApp/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middleware.CustomLogger())

	_routes.RegisterPathUser(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}
