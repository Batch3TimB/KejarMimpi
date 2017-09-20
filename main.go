package main

import (
	_ "github.com/kejarmimpi/routers"
	//"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))
	//orm.RegisterDataBase("default", "postgres", "user=postgres password=091289 host=127.0.0.1 dbname=kejarmimpi3")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
