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

// 路由响应格式化输出企业列表
func ReturnCompanyRespon(pageNum int, pageSize int, Company []model.Company_info) map[string]interface{} {

	var message_ = make(map[string]interface{})
	var data_all []interface{}
	fmt.Println("pageNum;", pageNum)
	fmt.Println("pageSize;", pageSize)
	for i, C := range Company {
		if i >= pageNum*pageSize || i < (pageNum-1)*pageSize {
			continue
		}

		m, err := StructToMap(C)
		if err != nil {
			message_["code"] = -1
			message_["data"] = make([]interface{}, 0)
			message_["msg"] = "接口调用获取失败"
			message_["page"] = map[string]int{
				"pageNum":  pageNum,
				"pageSize": 40,
				"total":    len(Company),
			}
			return message_
		} else {
			var data_one []map[string]interface{}
			for key, value := range m {
				var field = make(map[string]interface{})
				field["en_name"] = strings.ToLower(string(key[0])) + key[1:]
				field["value"] = value
				field["zh_name"] = nil
				data_one = append(data_one, field)
			}
			data_all = append(data_all, data_one)
		}
	}
	message_["code"] = 0
	var dataa = make(map[string]interface{})
	dataa["data"] = data_all
	message_["data"] = map[string]interface{}{"data": data_all}
	message_["msg"] = "成功"
	message_["page"] = map[string]int{
		"pageNum":  pageNum,
		"pageSize": 40,
		"total":    len(Company),
	}
	return message_
}

// 路由响应格式化输出人员列表
func ReturnPersonRespon(pageNum int, pageSize int, Person []model.Person_info) map[string]interface{} {

	fmt.Println("company_info:", len(Person))

	var message_ = make(map[string]interface{})
	var data_all []interface{}
	logger.Info("pageNum;", pageNum)
	logger.Info("pageSize;", pageSize)
	for i, C := range Person {

		fmt.Printf("i: %v", i)
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
	message_["code"] = 0
	var dataa = make(map[string]interface{})
	dataa["data"] = data_all
	message_["data"] = map[string]interface{}{"data": data_all}
	message_["msg"] = "成功"
	ReturnPage := map[string]int{
		"pageNum":  1,
		"pageSize": 40,
		"total":    len(Person),
	}
	message_["page"] = ReturnPage
	return message_
}

// 路由响应格式化输出企业
func CompanyOne(apply_num string, com model.Company_info) []interface{} {

	// fmt.Println("company_info:", len(Company))

	var data_all []interface{}
	var data_one []interface{}

	t := reflect.TypeOf(com)
	v := reflect.ValueOf(com)

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
	// var message = make(map[string]interface{})
	// message["code"] = 0
	// // var dataa = make(map[string]interface{})
	// // dataa["data"] = data_all
	// message["data"] = map[string]interface{}{
	// 	"companyInfoData": data_all,
	// 	"apply_num":       apply_num,
	// }
	// message["msg"] = "成功"

	// return message
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
