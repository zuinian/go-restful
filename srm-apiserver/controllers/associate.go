package controllers

import (
	"net/http"
	"github.com/astaxie/beego"
	"srm/srm-apiserver/models"
	"encoding/json"
	"fmt"
)

type AssociateController struct {
	beego.Controller
}

// 学籍修改-转学
// @router / [Put]
func (a *AssociateController) Put() {
	// 获取参数
	var associate models.Associate
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &associate)
	if err != nil {
		a.Data["json"] = ResponseError{"Invaild parameter param", 400,}
		a.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		a.ServeJSON()
		return
	}

	fmt.Print(associate)
	ok := models.AssociatePut(associate)
	if !ok {
		a.Data["json"] = ResponseError{"转学失败，学生或者学校信息有误", 400,}
		a.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		a.ServeJSON()
		return
	}
}
