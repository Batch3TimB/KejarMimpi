package controllers

import (
	"fmt"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/kejarmimpi/models"
	"github.com/astaxie/beego/orm"
)

// LoginController operations for authentication login user
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("LoginHandler", c.LoginHandler)
}

// Post ...
// @Title Post
// @Description authentication when user logging
// @Param	body		body 	models.Users	true		"body for Users content"
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) LoginHandler() {

	var data models.Users
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &data); err != nil {
        c.Ctx.Output.SetStatus(400)
        fmt.Println(c.Ctx.Input.RequestBody)
        c.Ctx.Output.Body([]byte("empty title"))
        return
    }
 
	res := struct{  
        Data []*models.Users
    }{}

    var cond *orm.Condition
    cond = orm.NewCondition()

    cond = cond.And("name", data.Name)
    cond = cond.And("password", data.Password)
    lim :=1

    var qs orm.QuerySeter
    qs = orm.NewOrm().QueryTable("users").SetCond(cond).Limit(lim)
    _,err := qs.All(&res.Data)
    if err != nil {
        c.Ctx.Output.SetStatus(400)
        c.Ctx.Output.Body([]byte(err.Error()))
        return
    }

	 

/*
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		fatal(err)
	}

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w)
*/
	c.ServeJSON()
}

