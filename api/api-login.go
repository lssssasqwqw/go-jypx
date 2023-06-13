package api

import (
	// "encoding/json"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	"net/http"
	// "reflect"
	// "strconv"
	// "strings"
	"time"
)

func SignIn(c *gin.Context) {
	data := map[string]interface{}{
		"codeIdentity": nil,
		"checkCode":    nil,
		"checkCodeImg": nil,
		"roleId":       "0",
		"roleName":     "超级管理员",
		"userId":       "0",
		"userName":     "admin",
		"token":        "NDc3YTgzMTYtNDg5Ny00ZWM3LWFiMmMtNzA4NGFjM2UxMTQw",
		"infoToken":    "",
	}

	signin_info := map[string]interface{}{
		"success":   "true",
		"code":      "000000",
		"message":   "ok",
		"timestamp": time.Now().Unix() * 1000,
		"data":      data,
		"msg":       "ok",
	}
	c.JSON(http.StatusOK, signin_info)
}

func QueryKey(c *gin.Context) {
	queryKey_info := map[string]interface{}{
		"code":    0,
		"message": "成功",
		"data":    "e90739f9-f191-4e4a-8f60-ce55804b4eb9",
		"page":    nil,
	}
	c.JSON(http.StatusOK, queryKey_info)
}

func SelectByCondition(c *gin.Context) {
	page := map[string]interface{}{
		"pageNum":  1,
		"pageSize": 10,
		"size":     1,
		"total":    0,
		"pages":    0,
	}
	condition := map[string]interface{}{
		"code":    0,
		"message": "成功",
		"data":    make([]int, 1),
		"page":    page,
	}
	c.JSON(http.StatusOK, condition)
}
