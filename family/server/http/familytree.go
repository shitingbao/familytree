package http

import (
	"familytree/family/lib/response"
	"familytree/family/model"
	"familytree/family/params"

	"github.com/gin-gonic/gin"
)

func FamilytreeCreate(ctx *gin.Context) {
	arg := &model.Member{}
	if err := ctx.Bind(&arg); err != nil {
		response.ApiFailed(ctx, err)
		return
	}
	err := Svc.FamilytreeCreate(arg)
	if err != nil {
		response.ApiFailed(ctx, err)
		return
	}
	response.ApiSuccess(ctx)
}

func FamilytreeUpdate(ctx *gin.Context) {
	arg := &params.ArgMemeber{}
	if err := ctx.Bind(&arg); err != nil {
		response.ApiFailed(ctx, err)
		return
	}
	err := Svc.FamilytreeUpdate(arg)
	if err != nil {
		response.ApiFailed(ctx, err)
		return
	}
	response.ApiSuccess(ctx)
}

func FamilytreeDelete(ctx *gin.Context) {
	arg := &params.ArgMemeber{}
	if err := ctx.Bind(&arg); err != nil {
		response.ApiFailed(ctx, err)
		return
	}
	err := Svc.FamilytreeDelete(arg)
	if err != nil {
		response.ApiFailed(ctx, err)
		return
	}
	response.ApiSuccess(ctx)
}

func FamilytreeList(ctx *gin.Context) {
	arg := &params.ArgMemeber{}
	if err := ctx.Bind(&arg); err != nil {
		response.ApiFailed(ctx, err)
		return
	}
	list, err := Svc.FamilytreeList(arg)
	if err != nil {
		response.ApiFailed(ctx, err)
		return
	}
	response.ApiSuccess(ctx, list)
}
