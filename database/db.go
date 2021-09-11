package database

import (
	"fmt"

	"github.com/rahulgarg03/ind21-rg-golang/src"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=myDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully Connected!!")
	}
	db.AutoMigrate(&src.User{})
	return db
}
