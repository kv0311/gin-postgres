package routes

import (
	"gin-demo/handler"

	"gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(routers *gin.Engine) {
	publicAPI := routers.Group("/", middleware.CheckToken)
	publicAPI.POST("/student/add", handler.AddStudent)
	publicAPI.DELETE("/student/delete", handler.DeleteStudent)
	publicAPI.POST("/student/update/name", handler.UpdateName)
	publicAPI.POST("/student/get", handler.GetStudent)
}
