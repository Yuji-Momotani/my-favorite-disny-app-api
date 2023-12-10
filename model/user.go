package model

import "time"

type User struct {
	// Gormが自動でusersテーブルに変換して作成
	ID        uint      `json:"id" gorm:"primaryKey"` //intのprimaryKeyのため、Gormが自動でautoincrementを付加
	Name      string    `json:"name" gorm:"unique; not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
