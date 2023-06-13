package pkg

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// 将结构体转换为 map
func StructToMap(obj interface{}) (map[string]interface{}, error) {
	// 将结构体转换为 JSON 字符串
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// 将 JSON 字符串解析为 map
	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取当前季度
func GetApplyNum() string {
	now := time.Now()
	nowDate := now.Format(Date)
	timeArray := strings.SplitN(nowDate, "-", 3)
	var r string
	s1, _ := strconv.Atoi(timeArray[1])
	if s1 < 4 {
		r = "4"
	} else if 4 <= s1 && s1 < 7 {
		r = "1"
	} else if 7 <= s1 && s1 < 10 {
		r = "2"
	} else {
		r = "3"
	}
	return timeArray[0] + "-" + r
}

func Strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))

	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
		// default:
		// 	newValue, _ := json.Marshal(value)
		// 	key = string(newValue)
	}
	return key
}
