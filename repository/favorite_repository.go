package repository

import (
	"errors"
	"my-favorite-disney-app-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// interface
type IFavoriteRepository interface {
	GetAllFavoritesByUser(favorites *[]model.Favorite, user_id uint) error
	CreateFavorites(favorite *model.Favorite) error
	UpdateFavorites(favorite *model.Favorite, character_id uint, user_id uint) error
	DeleteFavorites(character_id uint, user_id uint) error
}

// interfaceを実装するstruct
type favoriteRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewFavoriteRepository(db *gorm.DB) IFavoriteRepository {
	return &favoriteRepository{db: db}
}

// メソッドの定義
func (fr *favoriteRepository) GetAllFavoritesByUser(favorites *[]model.Favorite, user_id uint) error {
	if err := fr.db.Where("user_id = ?", user_id).Find(favorites).Error; err != nil {
		return err
	}
	return nil
}
func (fr *favoriteRepository) CreateFavorites(favorite *model.Favorite) error {
	if err := fr.db.Create(favorite).Error; err != nil {
		return err
	}
	return nil
}
func (fr *favoriteRepository) UpdateFavorites(favorite *model.Favorite, character_id uint, user_id uint) error {
	result := fr.db.Model(favorite).Clauses(clause.Returning{}).Where("character_id = ? AND user_id = ?", character_id, user_id).Update("evaluation", favorite.Evaluation)

	if err := result.Error; err != nil {
		return err
	}
	if rows := result.RowsAffected; rows < 1 {
		return errors.New("not found")
	}

	return nil
}
func (fr *favoriteRepository) DeleteFavorites(character_id uint, user_id uint) error {
	favorite := model.Favorite{}
	result := fr.db.Where("character_id = ? AND user_id = ?", character_id, user_id).Delete(&favorite)
	if err := result.Error; err != nil {
		return err
	}

	if rows := result.RowsAffected; rows < 1 {
		return errors.New("not found")
	}
	return nil
}
