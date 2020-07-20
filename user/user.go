package user

import (
	"github.com/gofiber/fiber"
)

func GetUsers(c *fiber.Ctx) {
	c.Send("All Users")
}

func GetUser(c *fiber.Ctx) {
	c.Send("Single User")
}

func NewUser(c *fiber.Ctx) {
	c.Send("New User")
}

func DeleteUser(c *fiber.Ctx) {
	c.Send("Delete User")
}
