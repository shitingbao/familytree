package http

import (
	"net/http"

	"familytree/family/params"

	"github.com/gin-gonic/gin"
)

func DyProjectCaseCreate(ctx *gin.Context) {
	arg := &params.ArgFamilyTree{}
	if err := ctx.Bind(&arg); err != nil {
		return
	}

	err := Svc.DyProjectCaseCreate()

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "ok",
	})
}
