// routes/routes.go
package routes

import (
	"authenserver/handlers"
	"authenserver/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Define user routes
	userRoutes := app.Group("/users")
	userRoutes.Post("/register", handlers.Register)
	userRoutes.Post("/login", handlers.Login)
	userRoutes.Post("/validate", handlers.ValidateToken)

	// Add more routes as needed
	bookRoute := app.Group("/book")
	bookRoute.Use(middleware.ValidateDomainMiddleware)
	bookRoute.Use(middleware.AuthMiddleware)
	bookRoute.Post("/insert", handlers.Insert)
}
