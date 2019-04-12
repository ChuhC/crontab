package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
)

var (
	 G_JobMgr *JobMgr
)

func init() {
	hosts := []string{beego.AppConfig.String("etcdHost")}
	config := clientv3.Config{
		Endpoints: hosts,
		DialTimeout: 5*time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil{
		logs.Error("ectd client init failed! err: %s ", err.Error())
	}
	G_JobMgr = &JobMgr{Client: client}
}
