package controller

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	authRouter(router.Group("/auth"))
	postRouter(router.Group("/post"))
}
