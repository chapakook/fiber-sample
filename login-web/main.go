package main

import (
	"fmt"
	"login-web/database"
	"login-web/env"
	"login-web/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env := env.Read()
	users := make(map[string]string)

	account := database.InitAccount(env["MYSQL_DRIVER"], env["MYSQL_SOURCE"])

	for _, a := range account {
		users[a.Id] = a.Pwd
	}

	config := middleware.BasicAuthConfigAllowUser(users)

	for id, pwd := range users {
		fmt.Println("ID -", id)
		fmt.Println("PWD -", pwd)
	}

	app := fiber.New()
	middleware.BasicAuth(app, config)
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Welcome") })

	app.Listen(":3000")
}
