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

func SignIn(c *fiber.Ctx) error {
    var userData models.User

    // Parse request body into userData
    if err := c.BodyParser(&userData); err != nil {
        return err
    }

    // Check if the user object is valid
    if userData.Username == "" || userData.Password == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Username and password are required",
        })
    }

    // Check if the user exists in the database
    var existingUser models.User
    if err := config.Db.Where("username = ?", userData.Username).First(&existingUser).Error; err != nil {
        if err !={
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Invalid username or password",
            })
        }
        return err
    }

    // Here you would typically check if the password matches for the existing user.
    // For this example, let's assume the logic is correct.
    
    return c.JSON(fiber.Map{
        "message": "User signed in successfully",
    })
}

