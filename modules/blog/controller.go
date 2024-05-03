package blog

import (
	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	root := router.Group("/blog")
	root.GET("", getAll)
	root.POST("", createOne)
	root.GET("/:id", getOne)
	root.PUT("/:id", updateOne)
	root.DELETE("/:id", deleteOne)
}
