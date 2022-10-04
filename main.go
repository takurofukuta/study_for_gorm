package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	const (
		// ユーザ名
		dbu = "go_study"
		// パスワード
		dbpw = "password"
		// ホスト名
		dbh = "127.0.0.1"
		// ポート
		dbp = "3306"
		// データベース名
		dbn = "go_study_db"
	)

	// DSN（Data Source Nane）の設定
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbu, dbpw, dbh, dbp, dbn)

	// DBへアクセス
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
}
