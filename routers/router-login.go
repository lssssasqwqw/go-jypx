package routers

import (
	"build/api"
	"github.com/gin-gonic/gin"
)

func Sign(r *gin.Engine) {
	S := r.Group("/auth")
	{
		S.POST("/user/signIn", api.SignIn)
	}
}

func Controller(r *gin.Engine) {
	S := r.Group("/controller")
	{
		S.GET("/user/queryKey", api.QueryKey)
		S.POST("/processMgt/selectByCondition", api.SelectByCondition)
	}
}
