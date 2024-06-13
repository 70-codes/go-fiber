package config

import (
	"fmt"
	"os"
	models "rest_api/Models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
	// // connection := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Connection failed")
	}
	DB = db
	fmt.Println("Connection successful")
	Automigrate(db)
}

func Automigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	)
}
