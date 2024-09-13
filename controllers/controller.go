package controllers

import (
	"am_blog/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var blog []models.Blog

	models.DB.Db.Find(&blog)

	return c.Status(fiber.StatusOK).JSON(blog)

}

func Create(c *fiber.Ctx) error {

	blog := new(models.Blog)

	if err := c.BodyParser(blog); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Create(&blog)

	return c.Status(fiber.StatusCreated).JSON(blog)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var blog models.Blog

	result := models.DB.Db.Where("id = ?", id).First(&blog)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "Blog ID not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.JSON(blog)
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")
	var updatedData models.Blog

	// Parse the body to get the updated user data
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if models.DB.Db.Where("id = ?", id).Updates(&updatedData).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID not found.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data has updated.",
	})

}

/* func Delete(c *fiber.Ctx) error {

	username := c.Params("username")

	// Delete the user record
	result := models.DB.Db.Where("username = ?", username).Delete(&models.Tryout{})
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).SendString("User deleted successfully")
} */
