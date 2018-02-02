// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"srm/srm-apiserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/schools",&controllers.SchoolController{},"post:Post")
	beego.Router("/api/v1/schools",&controllers.SchoolController{},"get:GetAll")
	beego.Router("/api/v1/schools/:schoolName/students",&controllers.SchoolController{},"get:StudentsBySchoolGet")

	beego.Router("/api/v1/students",&controllers.StudentController{},"post:Post")
	beego.Router("/api/v1/students/:studentName",&controllers.StudentController{},"put:Put")
	beego.Router("/api/v1/students/:studentName",&controllers.StudentController{},"get:Get")
	beego.Router("/api/v1/students/:studentName",&controllers.StudentController{},"delete:Delete")

	beego.Router("/api/v1/xueji",&controllers.AssociateController{},"put:Put")
}
