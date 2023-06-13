package api

import (
	"build/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"strconv"
	// "io/ioutil"
	"net/http"
	// "reflect"
	// "strconv"
	"gorm.io/gorm"
	// "strings"
)

type Data1 struct {
	DB *gorm.DB
}

type detail_pay_info model.Detail_pay_info

func (con Data1) TestInsert(c *gin.Context) {
	var FromData map[string]*string
	if err := c.ShouldBind(&FromData); err == nil {
	} else {
		fmt.Println("err:", err)
	}
	cardId := "421222xxxxyyss1221"
	checked := true
	result, _ := strconv.ParseBool(*FromData["isCreat"])
	fmt.Printf("result: %v\n", result)
	PayInfo := &model.Detail_pay_info{
		Id:          88888,
		D_apply_num: FromData["apply_num"],
		D_ID:        &cardId,
		Is_checked:  &checked,
		D_test:      3.0,
	}
	con.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "Id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"D_apply_num": PayInfo.D_apply_num}), // 主键重复，选择更新重复的列
	}).Session(&gorm.Session{SkipHooks: true}).Create(PayInfo)

	response := map[string]interface{}{
		"code":    0,
		"message": "成功",
		"data":    "e90739f9-f191-4e4a-8f60-ce55804b4eb9",
		"page":    nil,
	}
	c.JSON(http.StatusOK, response)
}

func (con Data1) TestSelect(c *gin.Context) {
	var result model.Person_info
	//con.DB.Joins("left join person_info on person_info.p_ID = detail_pay_info.d_ID").Where("d_test = ?", 20).Find(&result)
	con.DB.Where("p_ID = ?", "034046526308230044").Find(&result)
	fmt.Printf("result: %v\n", result)
	c.JSON(http.StatusOK, result)
}

func (con Data1) Index(c *gin.Context) {

	result := [...]map[string]string{{"name": "Runoob", "href": "single.html", "src": "images/work_16.jpg", "alt": "111", "title": "111", "num": "111"},
		{"name": "Google", "href": "single.html", "src": "images/work_16.jpg", "alt": "222", "title": "222", "num": "222"},
		{"name": "Taobao", "href": "single.html", "src": "images/work_16.jpg", "alt": "333", "title": "333", "num": "333"},
		{"name": "Taobao1", "href": "single.html", "src": "images/work_16.jpg", "alt": "333", "title": "333", "num": "333"},
		{"name": "Taobao2", "href": "single.html", "src": "images/work_16.jpg", "alt": "333", "title": "333", "num": "333"},
		{"name": "Taobao3", "href": "single.html", "src": "images/work_16.jpg", "alt": "333", "title": "333", "num": "333"},
		{"name": "Taobao4", "href": "single.html", "src": "images/work_16.jpg", "alt": "333", "title": "333", "num": "333"},
		{"name": "Taobao4", "href": "single.html", "src": "images/work_16.jpg", "alt": "333", "title": "333", "num": "333"}}

	fmt.Println(result)
	c.JSON(http.StatusOK, result)

}
