package api

type response struct {
	code int8
	data struct {
		data []interface{}
	}
	msg  string
	page struct {
		pageNum  int
		pageSize int
		total    int
	}
}

func (res *response) StruToMap() map[string]interface{} {
	msg := map[string]interface{}{
		"code": res.code,
		"data": map[string]interface{}{
			"data": res.data.data,
		},
		"msg": res.msg,
		"page": map[string]int{
			"pageNum":  res.page.pageNum,
			"pageSize": res.page.pageSize,
			"total":    res.page.total,
		},
	}
	return msg
}
