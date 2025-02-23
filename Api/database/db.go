package database

import (
	"log"

	"github.com/HugoCBB/Api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "root:hugo00028922@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatal(err, "Erro ao se conectar com o banco de dados")
	}

	DB.AutoMigrate(&models.Produtos{})
}
