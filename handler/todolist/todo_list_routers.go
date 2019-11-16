package todolist

import "github.com/gin-gonic/gin"

func TodoRouters(router *gin.RouterGroup) {
	router.GET("/todos", getTodos)
}
