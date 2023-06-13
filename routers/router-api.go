package routers

import (
	"build/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//22222222

func Select(r *gin.Engine, D *gorm.DB) {
	a := api.DataStruct{DB: D}
	S := r.Group("/api/v1")
	{
		S.POST("/companyList", a.CompanyList)

		S.GET("/checkDataIsWrong", a.CheckDataIsWrong)
		S.POST("/approvedList", a.ApprovedList)

		//【进入审核】
		S.POST("/calPersonCount", a.CalPersonCount)
		S.POST("/showAllFailPersonComment", a.ShowAllFailPersonComment)
		S.GET("/companyInfo", a.CompanyInfo)
		S.POST("/approvedCompany", a.ApprovedCompany)

		//【个人审核】
		S.GET("/personInfo", a.PersonInfo)
		S.GET("/checkRecord", a.CheckRecord)

		//【保存】金额保存
		S.POST("/approvedPersonInfo", a.ApprovedPersonInfo)
	}
}
