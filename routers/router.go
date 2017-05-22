// @APIVersion 1.0.0
// @Title log-service API
// @Description log-service only serve account register/delete/update/get
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"app-service/log-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/action",
			beego.NSInclude(
				&controllers.ActionController{},
			),
		),
		beego.NSNamespace("/log",
			beego.NSInclude(
				&controllers.LogController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
