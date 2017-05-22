package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["app-service/log-service/controllers:ActionController"] = append(beego.GlobalControllerRouter["app-service/log-service/controllers:ActionController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["app-service/log-service/controllers:ActionController"] = append(beego.GlobalControllerRouter["app-service/log-service/controllers:ActionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["app-service/log-service/controllers:LogController"] = append(beego.GlobalControllerRouter["app-service/log-service/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetPodLogById",
			Router: `/pod/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
