package userController

import (
	"github.com/gofiber/fiber/v2"
	"go-swaggo/models"
)

// HandleCreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param payload body models.UserModel true "UserModel"
// @Success 200 {string} string	"ok"
// @Router /user/create [post]
func HandleCreateUser(c *fiber.Ctx) error {
	var users models.UserModel
	err := c.BodyParser(&users)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "error",
			"data":    err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}
