package Controllers

import (
	"GoRest/Models"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetFoods ... Get all categories
func GetFoods(c *gin.Context) {
	id := c.Params.ByName("id")
	var food []Models.Food
	err := Models.GetAllFoodsByCategoryId(&food, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, food)
	}
}
