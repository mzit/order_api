package main

import (
	"fmt"
	"github.com/mzit/order_api/models"
	"github.com/mzit/order_api/pkg/logging"
	"github.com/mzit/order_api/pkg/setting"
	"github.com/mzit/order_api/routers"
	"net/http"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
