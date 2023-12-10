package controller

import (
	"my-favorite-disney-app-api/model"
	"my-favorite-disney-app-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// interface
type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	GetCsrf(c echo.Context) error
}

// interfaceを実装するstruct
type userController struct {
	uu usecase.IUserUsecase
}

// コンストラクタ
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu: uu}
}

// メソッド定義
func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	resUser, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, resUser)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	token, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   os.Getenv("API_DOMAIN"),
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true, //Postmanでテストするときはコメントアウト
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // フロントSPAのため、None
	}
	c.SetCookie(&cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) Logout(c echo.Context) error {
	cookie := http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Domain:  os.Getenv("API_DOMAIN"),
		Expires: time.Now(),
		// Secure:   true, //Postmanでテストするときはコメントアウト
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // フロントSPAのため、None
	}
	c.SetCookie(&cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) GetCsrf(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
