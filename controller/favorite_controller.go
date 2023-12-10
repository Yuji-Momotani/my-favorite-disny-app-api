package controller

import (
	"my-favorite-disney-app-api/model"
	"my-favorite-disney-app-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// interface
type IFavoriteController interface {
	GetAllFavorites(c echo.Context) error
	CreateFavorites(c echo.Context) error
	UpdateFavorites(c echo.Context) error
	DeleteFavorites(c echo.Context) error
}

// interfaceを実装するstruct
type favoriteController struct {
	fu usecase.IFavoriteUsecase
}

// コンストラクタ
func NewFavoriteController(fu usecase.IFavoriteUsecase) IFavoriteController {
	return &favoriteController{fu: fu}
}

func getUserIdFromJWT(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(float64)

	return uint(user_id)
}

// 実装部
func (fc *favoriteController) GetAllFavorites(c echo.Context) error {
	user_id := getUserIdFromJWT(c)

	// favorite := model.Favorite{}
	// if err := c.Bind(&favorite); err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }

	resFavorites, err := fc.fu.GetAllFavorites(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resFavorites)
}

func (fc *favoriteController) CreateFavorites(c echo.Context) error {
	favorite := model.Favorite{}
	user_id := getUserIdFromJWT(c)
	if err := c.Bind(&favorite); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}
	favorite.UserID = user_id
	newFavorite, err := fc.fu.CreateFavorites(favorite)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newFavorite)
}

func (fc *favoriteController) UpdateFavorites(c echo.Context) error {
	favorite := model.Favorite{}
	if err := c.Bind(&favorite); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user_id := getUserIdFromJWT(c)
	s_character_id := c.Param("characterId")
	character_id, _ := strconv.Atoi(s_character_id)

	updateFavorite := model.Favorite{
		CharacterID: uint(character_id),
		UserID:      user_id,
		Evaluation:  favorite.Evaluation,
	}
	resFavorite, err := fc.fu.UpdateFavorites(updateFavorite, uint(character_id), user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resFavorite)
}

func (fc *favoriteController) DeleteFavorites(c echo.Context) error {
	user_id := getUserIdFromJWT(c)
	s_character_id := c.Param("characterId")
	character_id, _ := strconv.Atoi(s_character_id)
	if err := fc.fu.DeleteFavorites(uint(character_id), user_id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
