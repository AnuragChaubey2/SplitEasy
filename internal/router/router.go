package router

import (
	"spliteasy/internal/auth"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.Register)
	}

	return r
}
