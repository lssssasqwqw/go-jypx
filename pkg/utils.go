package pkg

import (
	"build/logger"
	"build/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"strings"
	// "build/api"
)

const (
	Date     = "2006-01-02"
	DateTime = "2006-01-02 15:04:05"
)

func StrToFloat(str interface{}) *float64 {
	flo, _ := strconv.ParseFloat(Strval(str), 32)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", flo), 64)
	return &value
}

func StrToInt(str interface{}) *int64 {
	INT, _ := strconv.ParseInt(Strval(str), 10, 64)
	return &INT
}

func StrToBool(str interface{}) *bool {

	var value bool
	if fmt.Sprint(str) == "true" {
		value = true
	} else {
		value = false
	}
	fmt.Println("value", value)
	return &value
}

// 获取post请求所有的参数，返回map类型
func PostAllForm(c *gin.Context) map[string]interface{} {
	var tmpParams2 = make(map[string]interface{})
	err3 := c.ShouldBind(&tmpParams2)
	if err3 != nil {
		fmt.Println(err3)
		result := map[string]interface{}{
			"error": "暂未获取数据",
		}
		return result
	}
	return tmpParams2
}

// 响应格式化输出企业列表
func ReturnCompanyRespon(Company model.Company_info) (error, []map[string]interface{}) {

	m, err := StructToMap(Company)
	if err != nil {
		logger.Error("company结构体转map失败%v\n", err)
		return err, []map[string]interface{}{}
	} else {
		var data_one []map[string]interface{}
		for key, value := range m {
			var field = make(map[string]interface{})
			field["en_name"] = strings.ToLower(string(key[0])) + key[1:]
			field["value"] = value
			field["zh_name"] = nil
			data_one = append(data_one, field)
		}
		return nil, data_one
	}
}

// 响应格式化输出人员列表
func ReturnPersonRespon(person model.Person_info) (error, []map[string]interface{}) {

	m, err := StructToMap(person)
	if err != nil {
		logger.Error("person结构体转map失败v\n", err)
		return err, []map[string]interface{}{}
	} else {
		var data_one []map[string]interface{}
		for key, value := range m {
			var field = make(map[string]interface{})
			field["en_name"] = strings.ToLower(string(key[0])) + key[1:]
			field["value"] = value
			field["zh_name"] = nil
			data_one = append(data_one, field)
		}
		return nil, data_one
	}
}

// 路由响应格式化输出企业
func CompanyOne(apply_num string, com model.Company_info) []interface{} {

	// fmt.Println("company_info:", len(Company))

	var data_all []interface{}
	var data_one []interface{}

	m, err := StructToMap(com)
	if err != nil {
		logger.Error("解析出现错误%v\n", err)
	} else {
		for key, values := range m {
			var field = make(map[string]interface{})
			field["en_name"] = strings.ToLower(string(key[0])) + key[1:]
			field["value"] = values
			field["zh_name"] = nil
			data_one = append(data_one, field)
		}
	}
	data_all = append(data_all, data_one)
	return data_all
}

// 路由响应格式化输出人员
func ReturnOnePerson(C model.Person_info) []interface{} {

	var data_all []interface{}

	t := reflect.TypeOf(C)
	v := reflect.ValueOf(C)
	var data_one []map[string]interface{}
	for k := 0; k < t.NumField(); k++ {

		Lname := strings.ToLower(string(t.Field(k).Name[0])) + string(t.Field(k).Name[1:])

		var field = make(map[string]interface{})
		field["en_name"] = Lname
		field["value"] = v.Field(k).Interface()
		field["zh_name"] = nil
		data_one = append(data_one, field)
	}
	data_all = append(data_all, data_one)
	return data_all
}

// 路由响应格式化输出人员社保金额
func GetinsureMonery(Company []model.Detail_pay_info) []interface{} {

	var data_all []interface{}
	for _, C := range Company {

		t := reflect.TypeOf(C)
		v := reflect.ValueOf(C)
		var data_one []map[string]interface{}
		for k := 0; k < t.NumField(); k++ {
			Lname := strings.ToLower(string(t.Field(k).Name[0])) + string(t.Field(k).Name[1:])
			var field = make(map[string]interface{})
			field["en_name"] = Lname
			field["value"] = v.Field(k).Interface()
			field["zh_name"] = nil
			data_one = append(data_one, field)
		}
		data_all = append(data_all, data_one)
	}
	return data_all
}
