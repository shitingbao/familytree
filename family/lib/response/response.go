package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//错误码
const (
	CodeOk        = 10000 // 正确
	CodeSystemErr = 10001
)

const (
	MsgOk        = "ok" // 正确
	MsgSystemErr = "系统错误"
)

func ApiSuccess(ctx *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": CodeOk,
			"msg":  MsgOk,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": CodeOk,
		"msg":  MsgOk,
		"data": data[0],
	})
}

func ApiFailed(ctx *gin.Context, errInfo ...interface{}) {
	code := CodeSystemErr
	msg := MsgSystemErr
	if len(errInfo) > 0 {
		switch errMsg := errInfo[0].(type) {
		case string:
			msg = errMsg
		case error:
			msg = errMsg.Error()
		}
	}
	if len(errInfo) > 1 {
		code = errInfo[1].(int)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
