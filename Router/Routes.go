package Router

import (
	"GoRest/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}
	grp2 := r.Group("/category-api")
	{
		grp2.GET("categories", Controllers.GetCategories)
	}
	grp3 := r.Group("/food-api")
	{
		grp3.GET("foods/:id", Controllers.GetFoods)
	}

	return r
}
