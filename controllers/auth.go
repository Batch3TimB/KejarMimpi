package controllers

import (

	"github.com/astaxie/beego"
	"github.com/kejarmimpi/models"
)

type AuthController struct {
	beego.Controller

	Userinfo *models.Users
	IsLogin  bool
}

type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}


func (c *AuthController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}

func (c *AuthController) GetLogin() *models.Users {
	u := &models.Users{Id: c.GetSession("userinfo").(int)}
	u.Read()
	return u
}

func (c *AuthController) DelLogin() {
	c.DelSession("userinfo")
}

func (c *AuthController) SetLogin(user *models.Users) {
	c.SetSession("userinfo", user.Id)
}
