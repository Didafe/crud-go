package routes

import (
	"github.com/Didafe/crud-go/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserByID/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FinUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}
