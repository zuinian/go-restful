package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:AssociateController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:AssociateController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"Put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:SchoolController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:SchoolController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:SchoolController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:SchoolController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:SchoolController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:SchoolController"],
		beego.ControllerComments{
			Method: "StudentsBySchoolGet",
			Router: `/:schoolName/students`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:SchoolController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:SchoolController"],
		beego.ControllerComments{
			Method: "Associate",
			Router: `/associate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:StudentController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:StudentController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:StudentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:studentName`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:StudentController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:StudentController"],
		beego.ControllerComments{
			Method: "StudentGet",
			Router: `/:studentName`,
			AllowHTTPMethods: []string{"Get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:StudentController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:StudentController"],
		beego.ControllerComments{
			Method: "StudentDelete",
			Router: `/:studentName`,
			AllowHTTPMethods: []string{"Delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"] = append(beego.GlobalControllerRouter["srm/srm-apiserver/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
