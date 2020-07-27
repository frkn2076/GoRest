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
		// grp1.GET("user/:id", Controllers.GetUserByD)
		// grp1.PUT("user/:id", Controllers.UpdateUsr)
		// grp1.DELETE("user/:id", Controllers.DeleteUsr)
	}
	return r
}
