package controllers

import (
	"github.com/astaxie/beego"
	"srm/srm-apiserver/models"
	"encoding/json"
)

type StudentController struct {
	beego.Controller
}

// @router / [post]
func (s *StudentController) Post() {
	var student models.Student
	json.Unmarshal(s.Ctx.Input.RequestBody, &student)
	sid := models.AddStudent(student)
	s.Data["json"] = map[string]string{"sid": sid}
	s.ServeJSON()
}

// @router / [get]
func (s *StudentController) GetAll() {
	student := models.GetAllStudents()
	s.Data["json"] = student
	s.ServeJSON()
}

