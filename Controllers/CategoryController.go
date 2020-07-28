package Controllers

import (
	"GoRest/Models"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetCategories ... Get all categories
func GetCategories(c *gin.Context) {
	var category []Models.Category
	err := Models.GetAllCategories(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}
}
