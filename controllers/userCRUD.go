package controllers

import (
	"github.com/gofiber/fiber/v2"
	"weekOne/models"
	"weekOne/config"
)



func SignUp(c *fiber.Ctx) error {
    newUser := new(models.User)

    // Parse request body into newUser
    if err := c.BodyParser(newUser); err != nil {
        return err
    }

    // Check if the user object is valid
    if newUser.Username == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Username and password are required",
        })
    }

    // logic
    config.Db.Create(&newUser)
    return c.JSON(fiber.Map{
        "message": "User signed up successfully",
    })
}

