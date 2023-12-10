package validator

import (
	"my-favorite-disney-app-api/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Interface
type IFavoriteValidator interface {
	FavoriteValidator(favorite model.Favorite) error
}

// struct
type favoriteValidator struct{}

// コンストラクタ
func NewFavoriteValidator() IFavoriteValidator {
	return &favoriteValidator{}
}

// バリデーション処理
func (fv *favoriteValidator) FavoriteValidator(favorite model.Favorite) error {
	return validation.ValidateStruct(
		&favorite,
		validation.Field(
			&favorite.CharacterID,
			validation.Required.Error("CharacterID is Required"),
		),
		validation.Field(
			&favorite.UserID,
			validation.Required.Error("UserID is Required"),
		),
		validation.Field(
			&favorite.Evaluation,
			validation.Required.Error("Evaluation is Required"),
			validation.In(uint(1), uint(2), uint(3)).Error("Evaluation must be between 1 and 3"),
			validation.Min(uint(1)).Error("Evaluation must be between 1 and 3"),
			validation.Max(uint(3)).Error("Evaluation must be between 1 and 3"),
		),
	)
}
