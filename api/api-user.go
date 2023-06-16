package api

import (
	"build/logger"
	"build/model"
	"build/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	"net/http"
	// "reflect"
	// "strconv"
	"gorm.io/gorm"
	// "strings"
)

type DataStruct struct {
	DB *gorm.DB
}

func (con DataStruct) CompanyList(c *gin.Context) {
	var res response
	type company_ struct {
		Apply_season    string         `json:"apply_season" form:"apply_season"`
		C_person_result int            `json:"c_person_result" form:"c_person_result"`
		Page            map[string]int `json:"page" form:"page"`
		Token           string         `json:"token" form:"token"`
	}
	com := &company_{}
	if err := c.ShouldBind(&com); err == nil {
		fmt.Println("com", com)
	} else {
		fmt.Println("shouldbind_err:", err)
	}
	c_person_result := com.C_person_result
	pageSize := com.Page["pageSize"] //一页条数
	pageNum := com.Page["pageNum"]   //第几页

	apply_season := com.Apply_season
	if apply_season == "" {
		apply_season = pkg.GetApplyNum()
	}

	var Companys []model.Company_info
	if c_person_result == 2 {
		con.DB.Where("c_person_result is null and c_apply_num like ?", "%"+apply_season+"%").Find(&Companys)
	} else {
		con.DB.Where("c_person_result = ? and c_apply_num like ?", c_person_result, "%"+apply_season+"%").Find(&Companys)
	}

	var data_all []interface{}

	for i, C := range Companys {
		if i >= pageNum*pageSize || i < (pageNum-1)*pageSize {
			continue
		}
		err, data_one := pkg.ReturnCompanyRespon(C)
		if err == nil {
			data_all = append(data_all, data_one)
		} else {
			logger.Error("err:%v\n", err)
			res.code = -1
			res.msg = "失败"
			res.data.data = make([]interface{}, 0)
			c.JSON(http.StatusOK, res.StruToMap())
			return
		}
	}
	res.code = 0
	res.data.data = data_all
	res.page.pageNum = pageNum
	res.page.total = len(data_all)
	res.page.pageSize = 40
	res.msg = "成功"

	c.JSON(http.StatusOK, res.StruToMap())
}

// 检测数据
func (con DataStruct) CheckDataIsWrong(c *gin.Context) {
	// apply_num := c.Query("apply_num")

	data := map[string]interface{}{
		"code": 0,
		"data": nil,
		"msg":  "成功",
	}
	c.JSON(http.StatusOK, data)
}

// 获取企业下面成功，失败，预审人员详细数据
func (con DataStruct) ApprovedList(c *gin.Context) {
	var res response
	type Params struct {
		Apply_num string         `json:"apply_num" form:"apply_num"`
		List_type string         `json:"list_type" form:"list_type"`
		Page      map[string]int `json:"page" form:"page"`
	}
	params := &Params{}
	if err := c.ShouldBind(&params); err == nil {
		var person []model.Person_info

		logger.Error("params:%v", params)
		pageSize := params.Page["pageSize"]
		pageNum, ok := params.Page["pageNum"]
		if !ok {
			pageNum = 1
		}
		if params.List_type == "0" {
			con.DB.Where("p_c_apply_num = ? and p_person_result is null ", params.Apply_num).Find(&person)
		} else if params.List_type == "2" {
			con.DB.Where("p_c_apply_num = ? and p_person_result = 0", params.Apply_num).Find(&person)
		} else if params.List_type == "1" {
			con.DB.Where("p_c_apply_num = ? and p_person_result = 1", params.Apply_num).Find(&person)
		}
		var data_all []interface{}
		for i, p := range person {
			if i >= pageNum*pageSize || i < (pageNum-1)*pageSize {
				continue
			}
			err, data_one := pkg.ReturnPersonRespon(p)
			if err != nil {
				logger.Error("person结构体转map失败v\n", err)
				res.code = -1
				res.data.data = []interface{}{}
				res.msg = "失败"
				c.JSON(http.StatusOK, res.StruToMap())
				return
			} else {
				data_all = append(data_all, data_one)
			}
		}
		res.code = -1
		res.data.data = data_all
		res.msg = "成功"
		res.page.pageNum = pageNum
		res.page.total = len(data_all)
		res.page.pageSize = 40
		c.JSON(http.StatusOK, res.StruToMap())

	} else {
		res.code = -1
		res.data.data = []interface{}{}
		res.msg = "参数获取失败"
		c.JSON(http.StatusOK, res.StruToMap())
	}
}

// 获取企业下面成功，失败，预审数量
func (con DataStruct) CalPersonCount(c *gin.Context) {
	message_ := map[string]interface{}{
		"code": 0,
		"data": map[string]int{
			"failListCount":    0,
			"preditListCount":  5,
			"successListCount": 0,
		},
		"msg": "成功",
	}
	c.JSON(http.StatusOK, message_)
}

// 状态位，暂不清楚作用
func (con DataStruct) ShowAllFailPersonComment(c *gin.Context) {
	message_ := map[string]interface{}{
		"code": 0,
		"data": []string{},
		"msg":  "成功",
	}
	c.JSON(http.StatusOK, message_)
}

// 获取公司详细数据
func (con DataStruct) CompanyInfo(c *gin.Context) {
	apply_num := c.Query("apply_num")
	var com model.Company_info
	con.DB.Preload("Erson_info", func(db *gorm.DB) *gorm.DB {
		// 原生的sql语句, 这里的sno和cno是数据库中的字段，不是结构体中的字段
		return db.Where("P_ID=?", "004703522705283022")
	}).Where("c_apply_num = ?", apply_num).First(&com)

	data_all := pkg.CompanyOne(apply_num, com)
	// fmt.Printf("data_all: %v\n", data_all[1])
	var message = make(map[string]interface{})
	message["code"] = 0
	// var dataa = make(map[string]interface{})
	// dataa["data"] = data_all
	message["data"] = map[string]interface{}{
		"companyInfoData": data_all,
		"apply_num":       apply_num,
	}
	message["msg"] = "成功"

	c.JSON(http.StatusOK, message)
}

// //企业【审核通过】【审核不通过】
// func (con Data_) ApprovedCompany(c *gin.Context) {
// 	type CompanyCarry struct {
// 		Apply_num []string       `json:"apply_num"`
// 		Content   map[string]int `json:"content"`
// 	}

// 	Carry := &CompanyCarry{}
// 	if err := c.ShouldBind(Carry); err == nil {
// 	} else {
// 		fmt.Println("err;", err)
// 	}
// 	message_ := make(map[string]interface{})
// 	fmt.Printf("Carry: %v\n", Carry)
// 	c.JSON(http.StatusOK, message_)
// }

// 人员【审核通过】【审核不通过】
func (con DataStruct) ApprovedCompany(c *gin.Context) {
	type CompanyCarry struct {
		Apply_num []string       `json:"apply_num"`
		Content   map[string]int `json:"content"`
	}
	Carry := &CompanyCarry{}
	if err := c.ShouldBind(Carry); err == nil {
	} else {
		fmt.Println("err;", err)
	}
	message_ := make(map[string]interface{})
	fmt.Printf("Carry: %v\n", Carry)
	c.JSON(http.StatusOK, message_)
}

// 【个人审核】数据校验
func (con DataStruct) CheckRecord(c *gin.Context) {
	p_ID := c.Query("id")
	message_ := map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"id": p_ID,
			"recordInfoData": []interface{}{
				[]interface{}{
					map[string]interface{}{
						"en_name": "u_update_date",
						"value":   "2022-07-15 11:03:03",
						"zh_name": nil,
					},
					map[string]interface{}{
						"en_name": "u_result",
						"value":   true,
						"zh_name": nil,
					},
					map[string]interface{}{
						"en_name": "u_new_comment",
						"value":   nil,
						"zh_name": nil,
					},
					map[string]interface{}{
						"en_name": "u_name",
						"value":   "huangjl2",
						"zh_name": nil,
					},
				},
			},
			"msg": "成功",
		},
	}
	fmt.Printf("p_ID: %v\n", p_ID)
	c.JSON(http.StatusOK, message_)
}

// 【个人审核】数据查询返回
func (con DataStruct) PersonInfo(c *gin.Context) {
	var message = make(map[string]interface{})
	var monery []model.Detail_pay_info
	var per model.Person_info
	var com model.Company_info
	p_ID := c.Query("id")
	apply_num := c.Query("apply_num")
	// now_list := c.Query("now_list")
	con.DB.Where("c_apply_num = ?", apply_num).First(&com)
	con.DB.Where("p_c_apply_num = ? and p_ID = ? ", apply_num, p_ID).Find(&per)
	con.DB.Where("d_apply_num = ? and d_ID = ? and d_pay_month in ('202204', '202205', '202206') ", apply_num, p_ID).Find(&monery)

	companyInfoData := pkg.CompanyOne(apply_num, com)
	personInfoData := pkg.ReturnOnePerson(per)
	detailMonery := pkg.GetinsureMonery(monery)
	fmt.Println("detailMonery:", len(detailMonery))

	message["code"] = 0
	message["data"] = map[string]interface{}{
		"CompanyPreditResult":    map[string]interface{}{},
		"ID_list":                []string{p_ID},
		"PersonPreditResult":     map[string]interface{}{},
		"allowanceInfoData":      []string{},
		"companyInfoData":        companyInfoData,
		"personInfoData":         personInfoData,
		"company_location":       "广州市番禺区沙头街禺山西路329号16座1栋01单元101室",
		"company_person":         "李铭",
		"contract_register_date": "2021-07-22",
		"apply_num":              "0003131265112022-20",
		"hard_identify_date":     nil,
		"id":                     "340322199510184617",
		"degreeInfoData": []map[string]interface{}{{
			"admission_date":   "20140901",
			"card_ID":          "340322199510184617",
			"graduate_date":    "20180621",
			"graduate_num":     "*",
			"graduate_school":  "华南农业大学",
			"graduate_subject": "农业机械化及其自动化",
			"graduation":       "本科",
			"identifyCode":     "-",
			"name":             "沈昊",
			"study_type":       "普通全日制",
		}, {
			"admission_date":   "20140901",
			"card_ID":          "340322199510184617",
			"graduate_date":    "20180621",
			"graduate_num":     "*",
			"graduate_school":  "华南农业大学",
			"graduate_subject": "农业机械化及其自动化",
			"graduation":       "本科",
			"identifyCode":     "-",
			"study_type":       "普通全日制",
		}},
		"predit_count":         0,
		"recordInfoData":       []string{},
		"socialPayInfoData":    []string{},
		"start_contract_date":  "2021-07-12",
		"unemployInfoData":     []string{},
		"unemploy_end":         "2021-08-23",
		"unemploy_start":       "2021-06-29",
		"xwCompany_list":       []string{},
		"end_contract_date":    nil,
		"file_path_list":       []string{},
		"hardIdentifyInfoData": []string{},
		"businessInfoData":     []string{},
		"detailPayInfoData":    detailMonery,
	}
	// "degreeInfoData": []map[string]interface{},
	// messagyInfoData"] = companyInfoData
	// message["data"]["detailPayInfoData"] = detailPayInfoData
	message["msg"] = "成功"
	c.JSON(http.StatusOK, message)
}

func (con DataStruct) ApprovedPersonInfo(c *gin.Context) {
	type insure_ struct {
		Apply_num string                   `json:"apply_num" form:"apply_num"`
		Content   []map[string]interface{} `json:"content" form:"content"`
		Id        string                   `json:"id" form:"id"`
	}
	InsureArray := &insure_{}
	if err := c.ShouldBind(InsureArray); err == nil {
	} else {
		fmt.Println("err:", err)
	}

	fmt.Println("", InsureArray)
	id := InsureArray.Id
	content := InsureArray.Content
	apply_num := InsureArray.Apply_num
	D_old_insur := 0.0
	D_injury_insur := 0.0
	D_unemploy_insur := 0.0
	D_birth_insur := 0.0
	D_medical_insur := 0.0
	var D_normal_fund int64 = 0
	var month_all int64 = 0
	for _, v := range content {
		// var detail model.DetailPayInfo
		// con.DB.Where("d_ID = ? and d_apply_num and d_pay_month = ?", id, apply_num, model.Strval(v["pay_month"])).First(&detail)
		if fmt.Sprint(v["is_checked"]) == "true" {
			D_old_insur += *pkg.StrToFloat(v["old_insur"])
			D_injury_insur += *pkg.StrToFloat(v["injury_insur"])
			D_unemploy_insur += *pkg.StrToFloat(v["unemploy_insur"])
			D_birth_insur += *pkg.StrToFloat(v["birth_insur"])
			D_medical_insur += *pkg.StrToFloat(v["medical_insur"])
			D_normal_fund += *pkg.StrToInt(v["normal_fund"])
			month_all += 1
		}
		con.DB.Model(&model.Detail_pay_info{}).Where("d_ID = ? and d_apply_num = ? and d_pay_month = ?", id, apply_num, pkg.Strval(v["pay_month"])).Updates(model.Detail_pay_info{
			D_old_insur:      pkg.StrToFloat(v["old_insur"]),
			D_injury_insur:   pkg.StrToFloat(v["injury_insur"]),
			D_unemploy_insur: pkg.StrToFloat(v["unemploy_insur"]),
			D_birth_insur:    pkg.StrToFloat(v["birth_insur"]),
			D_medical_insur:  pkg.StrToFloat(v["medical_insur"]),
			D_normal_fund:    pkg.StrToInt(v["normal_fund"]),
			Is_checked:       pkg.StrToBool(v["is_checked"]),
		})
	}
	p_society_total := D_old_insur + D_injury_insur + D_unemploy_insur + D_medical_insur + D_birth_insur + float64(D_normal_fund)
	con.DB.Model(&model.Person_info{}).Where("p_ID = ? and p_c_apply_num = ? ", id, apply_num).Updates(model.Person_info{
		P_old_insur:      &D_old_insur,
		P_injury_insur:   &D_injury_insur,
		P_unemploy_insur: &D_unemploy_insur,
		P_medical_insur:  &D_medical_insur,
		P_birth_insur:    &D_birth_insur,
		P_normal_fund:    &D_normal_fund,
		P_apply_month:    &month_all,
		P_society_total:  &p_society_total,
	})

	message_ := map[string]interface{}{
		"code": 0,
		"data": nil,
		"msg":  "成功",
	}
	c.JSON(http.StatusOK, message_)
}
