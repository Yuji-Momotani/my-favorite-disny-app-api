package usecase

import (
	"my-favorite-disney-app-api/model"
	"my-favorite-disney-app-api/repository"
	"my-favorite-disney-app-api/validator"
)

// interface
type IFavoriteUsecase interface {
	GetAllFavorites(user_id uint) ([]model.FavoriteResponse, error)
	CreateFavorites(favorite model.Favorite) (model.FavoriteResponse, error)
	UpdateFavorites(favorite model.Favorite, character_id uint, user_id uint) (model.FavoriteResponse, error)
	DeleteFavorites(character_id uint, user_id uint) error
}

// interfaceを実装するstruct
type favoriteUsecase struct {
	fr repository.IFavoriteRepository
	fv validator.IFavoriteValidator
}

// コンストラクタ
func NewFavoriteUsecase(fr repository.IFavoriteRepository, fv validator.IFavoriteValidator) IFavoriteUsecase {
	return &favoriteUsecase{fr: fr, fv: fv}
}

// メソッドの定義
func (fu *favoriteUsecase) GetAllFavorites(user_id uint) ([]model.FavoriteResponse, error) {
	favorites := []model.Favorite{}
	if err := fu.fr.GetAllFavoritesByUser(&favorites, user_id); err != nil {
		return []model.FavoriteResponse{}, err
	}
	resFavorites := []model.FavoriteResponse{}
	for _, v := range favorites {
		resFavorite := model.FavoriteResponse{
			CharacterID: v.CharacterID,
			UserID:      v.UserID,
			Evaluation:  v.Evaluation,
		}
		resFavorites = append(resFavorites, resFavorite)
	}
	return resFavorites, nil
}
func (fu *favoriteUsecase) CreateFavorites(favorite model.Favorite) (model.FavoriteResponse, error) {
	if err := fu.fv.FavoriteValidator(favorite); err != nil {
		return model.FavoriteResponse{}, err
	}
	if err := fu.fr.CreateFavorites(&favorite); err != nil {
		return model.FavoriteResponse{}, err
	}
	resFavorite := model.FavoriteResponse{
		CharacterID: favorite.CharacterID,
		UserID:      favorite.UserID,
		Evaluation:  favorite.Evaluation,
	}
	return resFavorite, nil
}
func (fu *favoriteUsecase) UpdateFavorites(favorite model.Favorite, character_id uint, user_id uint) (model.FavoriteResponse, error) {
	if err := fu.fv.FavoriteValidator(favorite); err != nil {
		return model.FavoriteResponse{}, err
	}
	if err := fu.fr.UpdateFavorites(&favorite, character_id, user_id); err != nil {
		return model.FavoriteResponse{}, err
	}
	resFavorite := model.FavoriteResponse{
		CharacterID: favorite.CharacterID,
		UserID:      favorite.UserID,
		Evaluation:  favorite.Evaluation,
	}
	return resFavorite, nil
}
func (fu *favoriteUsecase) DeleteFavorites(character_id uint, user_id uint) error {
	if err := fu.fr.DeleteFavorites(character_id, user_id); err != nil {
		return err
	}
	return nil
}
