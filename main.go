package main

import (

	"os"

	_ "github.com/kejarmimpi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))
	//orm.RegisterDataBase("default", "postgres", "user=postgress password=password host=127.0.0.1 port=5432 dbname=mylocaldb")

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

