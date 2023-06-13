package routers

import (
	"build/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestApi(r *gin.Engine, D *gorm.DB) {

	S := r.Group("/api/v2")
	{
		S.POST("/insetr", api.Data1{DB: D}.TestInsert)
		S.GET("/select", api.Data1{DB: D}.TestSelect)
		S.GET("/index", api.Data1{DB: D}.Index)
	}
}
