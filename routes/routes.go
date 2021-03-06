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
	/*
	 * Register / Login
	 * Ungated Routes
	 */

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	/*
	 * Gated Routes
	 */
	app.Use(middleware.IsAuthenticated)

	// User Info
	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	// User Management
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)

	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	// Role Management
	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	// Permissions
	app.Get("/api/permissions", controllers.AllPermissions)

	// Products
	app.Get("/api/products", controllers.AllProducts)
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	// Image Upload
	app.Post("/api/upload", controllers.Upload)

	// Define Static URL for Uploads
	app.Static("/api/uploads", "./uploads")

	// Orders
	app.Get("/api/orders", controllers.AllOrders)
	app.Post("/api/export", controllers.Export)
	app.Get("/api/chart", controllers.Chart)
}
