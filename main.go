package main

import (
	"my-favorite-disney-app-api/controller"
	"my-favorite-disney-app-api/db"
	"my-favorite-disney-app-api/repository"
	"my-favorite-disney-app-api/router"
	"my-favorite-disney-app-api/usecase"
	"my-favorite-disney-app-api/validator"
)

func main() {
	connDB := db.NewDB()
	userValidator := validator.NewUserValidator()
	favoriteValidator := validator.NewFavoriteValidator()
	userRepository := repository.NewUserRepository(connDB)
	favoriteRepository := repository.NewFavoriteRepository(connDB)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	favoriteUsecase := usecase.NewFavoriteUsecase(favoriteRepository, favoriteValidator)
	userController := controller.NewUserController(userUsecase)
	favoriteController := controller.NewFavoriteController(favoriteUsecase)
	e := router.NewRouter(userController, favoriteController)
	e.Logger.Fatal(e.Start(":8080"))
}
