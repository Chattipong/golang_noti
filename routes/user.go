package routes

import (
	"golang_noti/controllers"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {

	userController := controllers.NewUserAPI()

	r.GET("/user/all", userController.GetAllUser)
	r.POST("/user/create", userController.CreateUser)
	r.PUT("/user/:id", userController.UpdateUserById)
	r.DELETE("/user/:id", userController.DeleteUserById)
}
