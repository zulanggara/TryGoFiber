package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulanggara/TryGoFiber/models"
	"github.com/zulanggara/TryGoFiber/repositories"
	"net/http"
)

var Users = []models.Users{}

func CreateNewUser(c *fiber.Ctx) error {
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result, err := repositories.CreateNewUser(user)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"responseCode":    200,
		"responseMessage": "Data created successfully!",
		"additionalInfo":  result,
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	result, err := repositories.GetAllUsers()
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"responseCode":    200,
		"responseMessage": "Data found!",
		"additionalInfo":  result,
	})
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := repositories.GetUserById(id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"responseCode":    200,
		"responseMessage": "Data found!",
		"additionalInfo":  result,
	})
}

func UpdateUserDataById(c *fiber.Ctx) error {
	id := c.Params("id")
	var updateduser models.Users

	if err := c.BodyParser(&updateduser); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	result, err := repositories.UpdateUserDataById(id, updateduser)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"responseCode":    200,
		"responseMessage": "Data changed successfully!",
		"additionalInfo":  result,
	})
}

func DeleteUserDataById(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.Users

	result, err := repositories.DeleteUserDataById(id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusNoContent).JSON(fiber.Map{
		"responseCode":    200,
		"responseMessage": "Data changed successfully!",
		"additionalInfo": fiber.Map{
			"result": result,
			"user":   user,
		},
	})

}

//-------------
//-------------
//-------------

var Languages = []models.ProgrammingLanguage{
	{Id: "1", Language: "C", Creator: "Dennis Ritchie"},
	{Id: "2", Language: "Java", Creator: "James Gosling"},
	{Id: "3", Language: "C++", Creator: " Bjarne Stroustrup"},
	{Id: "4", Language: "Python", Creator: "Guido van Rossum"},
}

func GetAllLanguagesData(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(Languages)
}

func CreateNewLanguageData(c *fiber.Ctx) error {
	var language models.ProgrammingLanguage

	if err := c.BodyParser(&language); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return err
	}
	Languages = append(Languages, language)
	return c.Status(http.StatusCreated).JSON(language)
}

func DeleteLanguageById(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, lang := range Languages {
		if lang.Id == id {
			Languages = append(Languages[:i], Languages[i+1:]...)
			break
		}
	}
	return c.Status(http.StatusNoContent).JSON(fiber.Map{"data": Languages})

}

func UpdateLanguageDataById(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedlanguage models.ProgrammingLanguage
	updatedlanguage.Id = id

	if err := c.BodyParser(&updatedlanguage); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return err
	}
	for i, lang := range Languages {
		if lang.Id == updatedlanguage.Id {
			Languages = append(Languages[:i], Languages[i+1:]...)
			Languages = append(Languages, updatedlanguage)
		}
	}
	return c.Status(http.StatusCreated).JSON(updatedlanguage)
}
