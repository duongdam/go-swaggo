package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "go-swaggo/docs"
	userController "go-swaggo/modules"
)

// @title GoFiber Example API
// @version 1.0
// @description Golang GoFiber swagger auto generate step by step by swaggo
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)             // default
	app.Post("/user/create", userController.HandleCreateUser) // create new user

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
