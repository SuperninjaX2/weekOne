package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"title": "weekone",
	}, "layout")
}
func Product(c *fiber.Ctx) error {

	return c.Render("partials/product-card",fiber.Map{})
}
func ProductPage(c *fiber.Ctx) error {

	return c.SendString("productPage")
}
