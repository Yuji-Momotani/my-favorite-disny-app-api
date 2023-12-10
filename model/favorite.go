package model

import "time"

type Favorite struct {
	// Gormが自動で「favoritesテーブル」として複数形の命名をしてくれる。
	ID          uint `json:"id" gorm:"primaryKey"` //int型のprimaryKeyでAutoIncrementを自動付加
	CharacterID uint `json:"character_id" gorm:"not null"`
	UserID      uint `json:"user_id" gorm:"not null"`
	User        User `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"` //モデル名 + IDでデフォルトで外部キーになるが、一応foreignKeyを指定
	Evaluation  uint `json:"evaluation" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FavoriteResponse struct {
	CharacterID uint `json:"character_id"`
	UserID      uint `json:"user_id"`
	Evaluation  uint `json:"evaluation"`
}
