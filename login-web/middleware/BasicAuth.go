package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	logrus "github.com/sirupsen/logrus"
)

var injectLogger = func(c *fiber.Ctx) error {
	username := c.Locals("username")
	logger := logrus.WithField("user", username)
	c.Locals("logger", logger)
	return c.Next()
}

func BasicAuth(app *fiber.App, config *basicauth.Config) {
	app.Use(basicauth.New(*config))
	app.Use(injectLogger)
}

func homeHandler(c *fiber.Ctx) {

}

func BasicAuthConfigAllowAll() *basicauth.Config {
	return &basicauth.Config{
		Authorizer: func(id, pwd string) bool {
			return true
		},
	}
}

func BasicAuthConfigAllowAdmin(admin map[string]string) *basicauth.Config {
	return &basicauth.Config{
		Users: admin,
	}
}

func BasicAuthConfigAllowUser(users map[string]string) *basicauth.Config {
	return &basicauth.Config{
		Users: users,
	}
}
