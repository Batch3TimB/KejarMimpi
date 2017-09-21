package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/kejarmimpi/lib"
)

// LoginController operations for authentication login user
type LoginController struct {
	AuthController
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
}


// Post ...
// @Title Post
// @Description authentication when user logging
// @Param	body		body 	models.Users	true		"body for Users content"
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Login() {
	flash := beego.NewFlash()
	email := c.GetString("email")

	c.Data["json"] =  c.Ctx.Input.RequestBody
	
	fmt.Printf("ini dibawah")

	password := c.GetString("password")

	user, err := lib.Authenticate(email, password)
	if err != nil || user.Id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Success logged in")
	flash.Store(&c.Controller)

	c.SetLogin(user)
}

