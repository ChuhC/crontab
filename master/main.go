package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/ChuhC/crontab/master/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// init logger
func  initLogger()  {
	config := make(map[string]interface{})
	config["filename"] = beego.AppConfig.String("logPath")

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
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// init logger
	initLogger()

	//beego.ErrorHandler("403", )

	beego.Run()
}
