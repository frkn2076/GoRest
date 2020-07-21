package category

import (
	"github.com/gofiber/fiber"
)

func GetCategories(c *fiber.Ctx){
	categories := [2]category{category{"1","a"}, category{"2","b"}}
	c.Send(categories)
}

type category struct {
	Category string
	Image string
}