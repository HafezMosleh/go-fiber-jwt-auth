package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

func main() {
    app := fiber.New()
    app.Post("/login", func(c *fiber.Ctx) error {
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user": "admin",
            "exp":  time.Now().Add(time.Hour * 72).Unix(),
        })
        t, err := token.SignedString([]byte("secret"))
        if err != nil { return c.SendStatus(fiber.StatusInternalServerError) }
        return c.JSON(fiber.Map{"token": t})
    })
    app.Listen(":3000")
}
