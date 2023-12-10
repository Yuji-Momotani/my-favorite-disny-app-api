package validator

import (
	"my-favorite-disney-app-api/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Interface
type IUserValidator interface {
	UserValidator(user model.User) error
}

// struct
type userValidator struct{}

// コンストラクタ
func NewUserValidator() IUserValidator {
	return &userValidator{}
}

// バリデーション処理
func (uv *userValidator) UserValidator(user model.User) error {
	return validation.ValidateStruct(
		&user,
		validation.Field(
			&user.Name,
			validation.Required.Error("Name is Required"),
			validation.RuneLength(4, 15).Error("Name must be between 4 and 15 character"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is Required"),
			validation.RuneLength(6, 15).Error("Password must be between 6 and 15 character"),
		),
	)
}
