package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:ArticlesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:CommentsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/kejarmimpi/controllers:UsersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
