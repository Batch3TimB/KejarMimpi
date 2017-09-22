package controllers

import (
	"fmt"
	"encoding/json"

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

type User struct {
    Email string
    Password string
}


// Post ...
// @Title Post
// @Description authentication when user logging
// @Param	body		body 	models.Users	true		"body for Users content"
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Login() {
	req := c.Ctx.Request     //in beego this.Ctx.Request points to the Http#Request
	p := make([]byte, req.ContentLength)    
	_, err := c.Ctx.Request.Body.Read(p)
	
	if err == nil {
	    var u User
	    err1 := json.Unmarshal(p, &u)
	    if err1 == nil {
	    	flash := beego.NewFlash()
	    	email := u.Email
	    	password := u.Password

	       user, err := lib.Authenticate(email, password)
			if err != nil || user.Id < 1 {
				flash.Warning(err.Error())
				flash.Store(&c.Controller)
				return
			}


			c.Data["json"] = user
	    } else {
	        fmt.Println("Unable to unmarshall the JSON request", err1);

	    }
	}

	c.ServeJSON()
}

