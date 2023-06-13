package model

import "gorm.io/gorm"

type Company_info struct {
	gorm.Model
	Id                    *int
	C_apply_num           *string `gorm:"primary_key"`
	C_register_num        *string
	C_name                *string
	C_social_num          *string
	C_size                *string
	C_bank                *string
	C_account             *string
	C_bank_card           *string
	C_charge_ID           *string `gorm:"column:c_charge_ID"`
	C_charge_name         *string // 法人名字
	C_contact_name        *string // 单位联系人
	C_contact_phone       *string // 联系电话
	C_apply_count         *int    // 申领人数
	C_hard_employ_count   *int    // 就业困难人数
	C_graduate_count      *int    // 高校毕业生人数
	C_army_count          *int    // 随军家属
	C_affect_person_count *int    // 受影响职工

	C_poor_count           *int     // 建档立卡贫困劳动人数
	C_total                *int     // 合计
	C_old_insur            *float32 // 养老保险金额
	C_unemploy_insur       *float32 // 失业保险金额
	C_injury_insur         *float32 // 工伤保险金额
	C_birth_insur          *float32 // 生育保险金额
	C_medical_insur        *float32 // 医疗保险金额
	C_social_total         *float32 // 社保补贴合计
	C_general_post_subsidy *float32 // 一般性岗位补贴
	C_special_fund_count   *int     // 专项资金人数
	C_unemploy_fund_count  *int     // 失业基金人数
	C_net_office           *bool    // 是否网办提交
	C_special_fund         *float32 // 专项资金金额
	C_unemploy_fund        *float32 // 失业基金金额
	C_success_count        *int     // 成功人数
	C_success_fund         *float32 // 成功总金额
	C_fail_count           *int     // 失败人数
	C_fail_fund            *float32 // 失败总金额
	// state 各种状态
	C_robot_result     *bool   // 机器人预审结果0为不通过1为通过)
	C_explain          *string // 预审说明
	C_detail           *string // 预审详情
	C_get_person_state *bool   `gorm:"column:c_get_person_state"` // 抓取员工数据状态0为未完成1为完成
	C_search_charge    *string // 法人名字结果(数据来源：市场监督局)
	C_search_location  *string // 企业地址结果(数据来源：市场监督局)
	C_search_name      *string // 企业名称结果
	C_person_result    *bool   // 人工审核结果(0为不通过1为通过)
	// C_date      db.Column(db.DateTime, nullable    True)  // 审核时间
	C_identify_person  *string // 审核人
	C_identify_comment *string // 审核说明
	C_apply_season     *string // 申报年季
	C_isnet            *bool   // 是否网办(0为否1为是)
	C_isprovince       *bool   // 是否省投保0为否1为是)
	C_isseason         *bool   // 是否本季度新增0为否1为是)
	C_season_add_count *int    // 本季度新增人数
	IsIdentify         *bool   `gorm:"column:isIdentify"` // 企业是否可审核
	C_audit_order      *int    // 审核顺序

	// // 额外参数
	C_extra_social_count *int  // 参保人数
	C_outputFile         *bool // 是否导出过表格
	C_isSendEmail        *bool // 是否发送过短信

	// 街道审核新增字段
	C_street_author *string // 街道审核人
	C_street_belong *string // 所属区街

	// // 时间
	// C_create_time      db.Column(db.DateTime, nullable    False, server_default    func.now())  // 创建时间
	// C_update_time      db.Column(TIMESTAMP, nullable    False)  // 更新时间

	Erson_info []Person_info `gorm:"foreignKey:P_c_apply_num;association_foreignkey:C_apply_num"`
}

type Person_info struct {
	Id                    *int
	P_c_apply_num         *string  //关联公司的申领编号（一对多）
	P_list_place          *string  //人员所在列表(0为成功，1为失败，2为待审核)
	P_ID_type             *string  `gorm:"column:p_ID_type"` //证件类型
	P_ID                  *string  `gorm:"column:p_ID"`      //证件号码
	P_name                *string  //名字
	P_sex                 *string  //性别
	P_age                 *int     //年龄
	P_type                *string  //人员类型
	P_hard_type           *string  //就业困难类别
	P_contract_start_date *string  //合同起始日期
	P_contract_end_date   *string  //合同结束日期
	P_old_insur           *float64 //养老保险
	P_injury_insur        *float64 //工伤保险
	P_unemploy_insur      *float64 //失业保险
	P_medical_insur       *float64 //医疗保险
	P_birth_insur         *float64 //生育保险

	P_society_total     *float64 //社保补贴合计
	P_normal_fund       *int64   //一般性岗位补贴
	P_apply_month       *int64   //申领月数
	P_iscarry           *bool    //是否结转(0为否1为是)
	P_unemploy_register *string  //失业登记退出去向
	P_verify_explain    *string  //审核说明
	P_pass_num          *string  //通行证号码
	P_social_num        *string  //个人社保号
	P_disabled_num      *string  //残疾人证书号码
	P_isgraduate        *string  //是否毕业2年内(0为否1为是)
	P_graduate_type     *string  //毕业生类别
	P_graduate_school   *string  //毕业院校
	P_graduate_date     *string  //毕业时间
	P_pass_date         *string  //考核合格时间
	P_isprovince        *string  //是否省投保为否1为是)
	P_help_type         *string  //资助类型
	P_reamark           *string  //备注
	P_new_comment       *string  //新增审核说明
	P_identify_date     *string  //经办日期
	P_local_area        *string  //经办机构

	//人员的失业登记日期、困难认定日期和就业登记日期
	P_start_unemploy_date    *string //失业始期 用于导出报表
	P_end_unemploy_date      *string //失业终期 用于导出报表
	P_hard_identify_date     *string //困难认定日期 用于导出报表
	P_start_contract_date    *string //合同始期 用于导出报表
	P_end_contract_date      *string //合同终期 用于导出报表
	P_contract_register_date *string //合同登记日期

	//state 各种状态
	P_detail_pay_state *bool   //是否抓取了补贴月份表格
	P_robot_result     *bool   //机器人预审结果False为不通过，True为通过)
	P_explain          *string //预审说明
	P_detail           *string //预审详情
	P_medical_check    *bool   //医保查询状态False为未完成，True为完成
	P_medical_update   *string //医保查询更新(0未更新1已更新）
	P_person_result    *bool   //人工审核结果(False为不通过，True为通过)
	// P_date = db.Column(db.DateTime, nullable=True, server_default=func.now())  //审核时间
	P_identify_person  *string //审核人
	P_identify_comment *string //审核说明

	P_social_company_identical *bool   //入职前参保单位一致
	P_file_path                *string //附件列表字符串

	P_social_state *bool //人员社保信息获取状态(默认为False)
	P_detail_info  *bool //人员详情信息状态(默认为False)
	// P_hasFile = db.Column(db.Boolean, server_default=text("False"))  //人员是否有附件
	//时间
	// P_create_time = db.Column(db.DateTime, nullable=False, server_default=func.now())  //创建时间
	// P_update_time = db.Column(TIMESTAMP, nullable=False)
}

type Back_record_info struct {
	Id              int    `gorm:"column:id"`
	Br_apply_season string // 申报年季
	// Br_apply_num       string // 申领编号
	Br_company         string // 企业名称
	Br_person_ID       string `gorm:"column:br_person_ID"` // 人员证件号码
	Br_person_name     string // 人员名字
	Br_fill_back       int    // 操作名称
	Br_detail          string // 操作详情
	Br_state           string // 操作状态
	Br_identify_person string // 审核人名字
	Br_cid             string // 关联id 用于定位个人

	// // 时间
	// br_create_date = db.Column(db.DateTime, nullable=False, server_default=func.now())  // 创建时间
	// br_update_date = db.Column(TIMESTAMP, nullable=True)  // 更新时间

}

type Detail_pay_info struct {
	Id int `gorm:"column:id"`
	// id = db.Column(db.Integer, primary_key=True, autoincrement=True, nullable=False)  //
	D_apply_num      *string
	D_ID             *string  `gorm:"column:d_ID"` // 证件号码(关联人员列表的证件号码
	D_pay_month      *string  // 投保日期
	D_old_insur      *float64 // 养老保险
	D_injury_insur   *float64 // 工伤保险
	D_unemploy_insur *float64 // 失业保险
	D_birth_insur    *float64 // 生育保险
	D_medical_insur  *float64 // 医疗保险
	D_normal_fund    *int64   // 一般性岗位补贴
	Is_checked       *bool    // 是否选中
	D_test           float64  `gorm:"default:18"`
}

// func (user *Person_info) person_info ()string {
// 	return "person_info"
// }
