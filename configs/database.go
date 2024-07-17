package configs

import (
	"fmt"
	"github.com/zulanggara/TryGoFiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	env := NewEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	DB.AutoMigrate(&models.ProgrammingLanguage{})
	DB.AutoMigrate(&models.Users{})
}
