package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/sudachi0114/go-mvc/models"
)

func main() {
	// db, _ := gorm.Open("mysql", "root:@/gorm?charset=uft8&parseTime=True") // DB に接続
	// db.CreateTable(&models.User{}) // User model に対応するテーブルを作成
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベースの接続に問題が発生しました。")
	}
	db.CreateTable(&models.User{})

	defer db.Close() // DB との接続を閉じる
}
