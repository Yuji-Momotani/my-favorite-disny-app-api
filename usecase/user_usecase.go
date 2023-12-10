package usecase

import (
	"my-favorite-disney-app-api/model"
	"my-favorite-disney-app-api/repository"
	"my-favorite-disney-app-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Login(user model.User) (string, error)
	SignUp(user model.User) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur: ur, uv: uv}
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidator(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByUserName(&storedUser, user.Name); err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidator(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		Name:     user.Name,
		Password: string(hash),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, nil
	}
	resUser := model.UserResponse{
		ID:   newUser.ID,
		Name: newUser.Name,
	}
	return resUser, nil
}
