package controllers

import (
	"app-service/log-service/models"
	"app-service/log-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Action
type ActionController struct {
	beego.Controller
}

// @Title Create
// @Description create action
// @Param	body		body 	models.Action	true		"body for action content"
// @Success 200 {object} models.Response
// @Failure 403 body is empty
// @router / [post]
func (this *ActionController) Create() {
	var err error
	var action model.Action
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &action)
	if err == nil {
		var svc service.ActionService
		err = svc.Create(&action)
		if err == nil {

			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = ""

		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetAll
// @Description get all user's actions
// @Param	userid		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @router /:userId [get]
func (this *ActionController) GetAll() {
	var err error
	var response models.Response

	var userId int64
	userId, err = this.GetInt64(":userId")
	beego.Debug("GetAllActions", userId)
	if userId > 0 && err == nil {
		var svc service.ActionService
		var actions []*model.Action
		var result []byte
		actions, err = svc.GetAll(userId)
		if err == nil {
			result, err = json.Marshal(actions)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
