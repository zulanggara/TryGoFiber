package configs

import (
	"github.com/zulanggara/TryGoFiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/rest_api_go?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	DB.AutoMigrate(&models.ProgrammingLanguage{})
	DB.AutoMigrate(&models.Users{})
}
