package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ChuhC/crontab/worker/mgr"
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego/logs"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	flag.StringVar(&confFile, "config", "src/github.com/Chuhc/crontab/worker/conf/worker.conf", "指定配置文件")
	flag.Parse()
}

// init logger
func initLogger(fileName string) {
	config := make(map[string]interface{})
	config["filename"] = fileName

	// map 转 json
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
		return
	}
	logs.Debug(string(configStr))
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.SetLogFuncCall(true)
}

func main() {

	initArgs()

	cfg, err := goconfig.LoadConfigFile(confFile)
	if err != nil {
		panic("load configure file error")
	}
	etcdHost, _ := cfg.GetValue("etcd", "host")
	etcdPort, _ := cfg.Int("etcd", "port")
	logPath, _ := cfg.GetValue("log", "path")

	// init logger
	initLogger(logPath)

	// init job mgr
	mgr.InitJobMgr(etcdHost, etcdPort)

	// catch ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	for {
		select {
		case <-c:
			return
		default:
			time.Sleep(time.Second)
		}
	}

}
