## Golang gofiber swagger auto generate step by step by swaggo

Follow https://github.com/swaggo/swag

Step 1: Install swaggo

```bash
	go install github.com/swaggo/swag/cmd/swag@latest
```

Step 2: Install gofiber swagger

```bash
	go get -u github.com/gofiber/swagger
```

Step 3: Add comments to your API source code

Example your go.mod

```bash
module go_api

go 1.20
```

Example your main.go

Notice the import ->  _ "go_api/docs", and change info of // @title... if you want.

```bash
package main

import (
	"github.com/gofiber/swagger"
	"github.com/gofiber/fiber/v2"

	_ "go_api/docs"
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

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Post("/user/create", userController.HandlerCreateUser) // create new user
	
	app.Listen(":8080")
}
```

Example your function create user

```bash
package userController

import (...)

// HandlerCreateUser godoc
//
//	@Summary		Create new user
//	@Description	Create new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		userModel.UserModel	true	"UserModel"
//	@Success		200		{string}  string  "OK"
//	@Failure		400		{string}  error  "Bad Request"
//	@Router			/user/create [post]
func HandlerCreateUser(c *fiber.Ctx) error {
...
}

# Notice: Remember line "//	@Router			/user/create [post]" and line "func HandlerCreateUser(c *fiber.Ctx) error {" do not have a newline.
```

-> Done step 3

Step 4: Run the Swag in your Go project root folder which contains main.go file, Swag will parse comments and generate
required files(docs folder and docs/doc.go).

 ```bash
 swag init
 ```

Step 5: Run go project and move to http://localhost:8080/swagger/index.html

 ```bash
 go run .
 ```

Thanks for reading