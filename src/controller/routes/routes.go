package routes

import (
	user_controller "github.com/RafaDRF/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/userById/:userId", user_controller.FindUserById)
	r.GET("/userByEmail/:userEmail", user_controller.FindUserByEmail)
	r.POST("/user", user_controller.CreateUser)
	r.PUT("/user/:userId", user_controller.UpdateUser)
	r.DELETE("/user/:userId", user_controller.DeleteUser)
}
