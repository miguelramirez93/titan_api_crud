package main

import (
	_ "titan_api_crud/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
}


func main() {
	orm.Debug = true
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
