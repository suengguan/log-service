package service

import (
	"fmt"
	"strconv"

	"model"
	"utility/fileoperator"

	daoApi "api/dao_service"

	"github.com/astaxie/beego"
)

type LogService struct {
}

func (this *LogService) GetPodLogById(id int64) (*model.PodLog, error) {
	var err error
	var podLog model.PodLog

	// get pod
	beego.Debug("->get pod")
	var pod *model.Pod
	pod, err = daoApi.BussinessDaoApi.GetPodById(id)
	if err != nil {
		beego.Debug("get pod failed")
		return nil, err
	}

	// get module
	beego.Debug("->get module")
	var module *model.Module
	module, err = daoApi.BussinessDaoApi.GetModuleById(pod.Module.Id)
	if err != nil {
		beego.Debug("get module failed")
		return nil, err
	}

	// get job
	beego.Debug("->get job")
	var job *model.Job
	job, err = daoApi.BussinessDaoApi.GetJobById(module.Job.Id)
	if err != nil {
		beego.Debug("get job failed")
		return nil, err
	}

	// get project
	beego.Debug("->get porject")
	var project *model.Project
	project, err = daoApi.BussinessDaoApi.GetProjectById(job.Project.Id)
	if err != nil {
		beego.Debug("get project failed")
		return nil, err
	}

	// get user
	beego.Debug("->get user")
	var user *model.User
	user, err = daoApi.UserDaoApi.GetById(project.User.Id)
	if err != nil {
		beego.Debug("get user failed")
		return nil, err
	}

	// get log
	beego.Debug("->get log")
	podLog.ProjectId = project.Id
	podLog.ProjectName = project.Name
	podLog.JobId = job.Id
	podLog.JobName = job.Name
	podLog.ModuleId = module.Id
	podLog.ModuleName = module.Name
	podLog.PodId = pod.Id
	podLog.PodName = pod.Name

	var fn string
	var cfg = beego.AppConfig
	// todo
	// if config.DEBUG_ONLY {
	// 	fn = "C:/Works/PME2017/solution/PMEServer/workspace/debug.log"
	// } else {
	// 	fn = config.WORKSPACE_PATH
	// 	// pod path
	// 	fn += "/" + user.Name
	// 	fn += "/" + project.Name + "-" + strconv.FormatInt(project.Id, 36)
	// 	fn += "/" + job.Name + "-" + strconv.FormatInt(job.Id, 36)
	// 	fn += "/" + module.Name + "-" + strconv.FormatInt(module.Id, 36)
	// 	fn += "/" + pod.Name + "-" + strconv.FormatInt(pod.Id, 36) + ".log"
	// }

	if cfg.String("runmode") == "dev" {
		fn = "C:/Works/PME2017/solution/PMEServer/workspace/debug.log"
	} else {
		fn = cfg.String("workspace")
		// pod path
		fn += "/" + user.Name
		fn += "/" + project.Name + "-" + strconv.FormatInt(project.Id, 36)
		fn += "/" + job.Name + "-" + strconv.FormatInt(job.Id, 36)
		fn += "/" + module.Name + "-" + strconv.FormatInt(module.Id, 36)
		fn += "/" + pod.Name + "-" + strconv.FormatInt(pod.Id, 36) + ".log"
	}

	// get log
	var logContent string
	logContent, err = fileoperator.Read(fn)

	if err != nil {
		beego.Debug("get pod log failed! reason : ", err)
		err = fmt.Errorf("%s", "get pod log failed! reason :"+err.Error())
		return nil, err
	} else {
		podLog.LogContent = logContent
	}

	beego.Debug("result:ok")

	return &podLog, err
}
