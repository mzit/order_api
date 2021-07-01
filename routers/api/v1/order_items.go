package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mzit/order_api/models"
	"github.com/mzit/order_api/pkg/app"
	"github.com/mzit/order_api/pkg/e"
	"github.com/unknwon/com"
	"net/http"
)

func GetOrderItems(ctx *gin.Context) {
	appG := app.Gin{
		C: ctx,
	}
	code := e.INVALID_PARAMS
	orderSn := com.ToStr(ctx.Param("order_sn"))
	var data interface{}
	if models.ExistOrderGoodsByOrderSn(orderSn) {
		data = models.GetOrderGoodsByOrderSn(orderSn)
		code = e.SUCCESS
	} else {
		code = e.NOT_FOUND
	}
	appG.Response(http.StatusOK, code, data)
}
