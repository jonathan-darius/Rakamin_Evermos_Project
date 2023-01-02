package router

import "github.com/gin-gonic/gin"

func AuthRouter(router *gin.RouterGroup) {
	group := router.Group("/auth")

	group.POST("/register")
	group.POST("/login")
}
