package router

import (
	"ToDo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	todogroup := r.Group("v1")
	{
		//待办事项
		//增加
		todogroup.POST("/todo", controller.CreateTodo)
		//查看
		//查看所有
		todogroup.GET("/todo", controller.GetTodoList)

		//修改
		todogroup.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		todogroup.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
