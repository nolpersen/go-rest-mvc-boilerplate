package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nolpersen/src/controllers"
)

func Routes() {
	route := gin.Default()
	route.POST("/todo", controllers.Store)
	route.GET("/todos", controllers.Index)
	route.PUT("/todo/update/:id", controllers.Update)
	route.POST("/todo/delete/:id", controllers.Delete)

	route.Run()
}
