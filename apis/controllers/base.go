package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type baseController struct{}

func NewBaseController() *baseController {
	return &baseController{}
}

func (bc *baseController) GetParamTypeNumber(ctx *gin.Context, param string) (int64, bool) {
	paramValue := ctx.Param(param)
	if paramValue == "" {
		return 0, false
	}

	num, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil {
		return 0, false
	}

	return num, true
}

func (bc *baseController) Bind(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return false
	}
	return true
}

func (bc *baseController) GetUserID(ctx *gin.Context) (int64, bool) {
	value, ok := ctx.Get("userId")
	if !ok {
		return 0, false
	}

	userID, ok := value.(int64)
	if !ok {
		return 0, false
	}

	return userID, true
}
