package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//配置的获取
	c.Ctx.WriteString("appname:"+beego.AppConfig.String("appname")+"\n"+
		"httpport:"+beego.AppConfig.String("httpport")+"\n"+
		"runmode:"+beego.AppConfig.String("runmode")+"\n"+
		"AppName:"+beego.BConfig.AppName+"\n")
		// "AppConfigPath:"+beego.AppConfigPath+"\n"
		
	
	//类型转换
	hp := strconv.Itoa(3)
	c.Ctx.WriteString("hp:"+hp)
	
	//日志输出与日志级别设置
	beego.Trace("I'm test beego Trace log")
	beego.Info("Info log Context")
	// beego.SetLevel(beego.LevelInfo)
	beego.Trace("222 I'm test beego Trace log")//
	beego.Info("222 Info log Context")
		
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
