// Routing
// Magic things happen here!
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scopehs/tutorial/controllers"
)

// Setup Route
func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/users", controllers.User)
}
