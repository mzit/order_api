package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mzit/order_api/models"
	"github.com/mzit/order_api/pkg/e"
	"github.com/unknwon/com"
	"net/http"
)

func GetOrderItems(ctx *gin.Context) {
	code := e.INVALID_PARAMS
	orderSn := com.ToStr(ctx.Param("order_sn"))
	var data interface{}
	if models.ExistOrderGoodsByOrderSn(orderSn) {
		data = models.GetOrderGoodsByOrderSn(orderSn)
		code = e.SUCCESS
	} else {
		code = e.NOT_FOUND
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
