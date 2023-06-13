package main

import (
	"build/config"
	"build/model"
	"build/routers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

//asdadasdafsafdfsdfs

func main() {
	DB = config.SqlConnet()
	DB.AutoMigrate(&model.Detail_pay_info{}) //数据库字段更新

	r := gin.Default()
	r.Use(Cors())

	routers.Sign(r)
	routers.Controller(r)
	routers.Select(r, DB)
	routers.TestApi(r, DB) //测试接口
	r.Run("localhost:5000")
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token, ddddd")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		context.Next()
	}
}
