package database

import (
	"api-go/handler/structures"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	Configs := "root:@tcp(127.0.0.1:3306)/api-book?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(Configs), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.Debug().AutoMigrate(&structures.Book{})

	return DB, nil
}
