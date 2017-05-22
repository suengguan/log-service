package controllers

import (
	"app-service/log-service/models"
	"app-service/log-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Log
type LogController struct {
	beego.Controller
}

// @Title GetPodLogById
// @Description get pod log by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @Failure 403 :id is empty
// @router /pod/:id [get]
func (this *LogController) GetPodLogById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	//beego.Debug("GetPodLogById", id)
	if id > 0 && err == nil {
		var svc service.LogService
		var podLog *model.PodLog
		var result []byte
		podLog, err = svc.GetPodLogById(id)
		if err == nil {
			result, err = json.Marshal(podLog)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "pod id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
