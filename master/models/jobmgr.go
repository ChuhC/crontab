package models

import (
	"context"
	"encoding/json"

	"github.com/ChuhC/crontab/common"
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
)

var (
)

// job manager
type JobMgr struct {
	Client *clientv3.Client // etcd client
}

func (j *JobMgr) SaveJob(job common.Job)(oldJob *common.Job, err error)  {
	key := "/cron/jobs/" + job.Name
	var jobValue []byte
	jobValue, err = json.Marshal(job)
	if err != nil{
		return
	}

	logs.Info("#############", string(jobValue), key)
	var putResp *clientv3.PutResponse
	getResp, err := j.Client.Get(context.TODO(), "test")
	if err != nil{
		println(err)
	}
	logs.Info(getResp)
	// save job to etcd
	putResp, err = j.Client.Put(context.TODO(), key, string(jobValue), clientv3.WithPrevKV())
	if err != nil{
		println("+++++++++++++++", err)
		logs.Error(err)
		return nil, err
	}

	// return old job if already exist
	if putResp.PrevKv != nil{
		var job = common.Job{}
		json.Unmarshal(putResp.PrevKv.Value, &job)
		oldJob = &job
		return oldJob, nil
	}
	return
}

func (j *JobMgr) DelJob(){

}

func (j *JobMgr) StopJob(){

}

func (j *JobMgr) GetAll(){

}

