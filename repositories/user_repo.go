package repositories

import (
	"github.com/zulanggara/TryGoFiber/configs"
	"github.com/zulanggara/TryGoFiber/models"
)

func CreateNewUser(user models.Users) (models.Users, error) {

	result := configs.DB.Create(&user)
	if result.Error != nil {
		return models.Users{}, result.Error
	}
	return user, nil
}

func GetAllUsers() ([]models.Users, error) {
	var user []models.Users

	result := configs.DB.Find(&user)
	if result.Error != nil {
		//return result.Error
		return []models.Users{}, result.Error
	}
	return user, nil
}

func GetUserById(id string) (models.Users, error) {
	var user models.Users

	result := configs.DB.Find(&user, "id = ?", id)
	if result.Error != nil {
		return models.Users{}, result.Error
	}
	return user, nil
}

func UpdateUserDataById(id string, user models.Users) (models.Users, error) {
	result := configs.DB.Model(&user).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return models.Users{}, result.Error
	}
	return user, nil
}

func DeleteUserDataById(id string) (models.Users, error) {
	var user models.Users
	result := configs.DB.Model(&user).Delete("id = ?", id)
	if result.Error != nil {
		return models.Users{}, result.Error
	}
	return user, nil
}
