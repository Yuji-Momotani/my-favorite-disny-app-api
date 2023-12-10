package main

import (
	"fmt"
	"my-favorite-disney-app-api/db"
	"my-favorite-disney-app-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	// 例：userとtaskテーブルを作成したい場合
	dbConn.AutoMigrate(&model.User{}, &model.Favorite{}) //作成したいモデルのstructを0値で引数に渡す
}
