package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	
	
	"weekOne/Routes"
	"weekOne/config"
)
func init() {
  config.Database()
}
func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
func test(){
	1+1
}
