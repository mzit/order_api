package v1

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mzit/order_api/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/mzit/order_api/models"
	"github.com/mzit/order_api/pkg/app"
	"github.com/mzit/order_api/pkg/e"
	"github.com/mzit/order_api/pkg/setting"
	"github.com/mzit/order_api/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

// GetOrders godoc example todo edit
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Header 200 {string} Token "qwerty"
// @Router /accounts/{id} [get]
func GetOrders(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	storeId := ctx.Query("store_id")

	where := make(map[string]interface{})
	data := make(map[string]interface{})

	if storeId != "" {
		where["store_id"] = storeId
	}

	var platform int = 1
	if arg := ctx.Query("state"); arg != "" {
		platform = com.StrTo(arg).MustInt()
		where["platform"] = platform
	}

	code := e.SUCCESS

	data["lists"] = models.GetOrders(util.GetOffset(ctx), setting.AppSetting.PageSize, where)
	data["total"] = models.GetOrderTotal(where)

	appG.Response(http.StatusOK, code, data)
}

func GetOrder(ctx *gin.Context) {

}
