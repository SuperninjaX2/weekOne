package routes

import (
	"github.com/gofiber/fiber/v2"
	"weekOne/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.Home)
	app.Get("/partials/product", controllers.Product)
	app.Get("/producr-page", controllers.ProductPage)
	app.Post("/authenticate/signup",controllers.SignUp)
}
