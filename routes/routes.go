// Routing
// Magic things happen here!
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scopehs/tutorial/controllers"
	"github.com/scopehs/tutorial/middleware"
)

// Setup Route
func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)

}
