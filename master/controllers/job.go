package controllers

import (
	"fmt"

	"github.com/ChuhC/crontab/common"
	"github.com/ChuhC/crontab/master/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)


// job controller
type  JobController struct {
	beego.Controller
}


// @Title Save
// @Description save job to etcd
// @Param	name     query 	    string	true        "crontab task name"
// @Param	crontab  query 	    string	true        "crontab expression"
// @Param	command  query	    string	true        "crontab command"
// @Success 200 {string} job save success
// @Failure 403 get no job
// @router /save [post]
func (j *JobController) Save(){

	job := common.Job{}
	if err := j.ParseForm(&job); err != nil{
		logs.Error("save job error")
	}
	oldJob, err := models.G_JobMgr.SaveJob(job)
	if err != nil{
		logs.Error("save job error")
	}
	if oldJob != nil{
		logs.Debug("job already exist. oldJob: %s", oldJob)
	}
	j.Ctx.WriteString(fmt.Sprintf("save job success! name: %s, cronExpr: %s, command: %s", job.Name,
		job.CronExpr, job.Command) )
}

