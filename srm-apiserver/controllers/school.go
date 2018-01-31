package controllers

import (
	"github.com/astaxie/beego"
	"srm/srm-apiserver/models"
	"encoding/json"
	"fmt"
)

// Operations about object
type SchoolController struct {
	beego.Controller
}

// @router / [post]
func (s *SchoolController) Post() {
	var school models.School
	json.Unmarshal(s.Ctx.Input.RequestBody, &school)
	sid := models.AddSchool(school)
	s.Data["json"] = map[string]string{"sid": sid}
	s.ServeJSON()
}

// @router / [get]
func (s *SchoolController) GetAll() {
	schools := models.GetAllSchools()
	s.Data["json"] = schools
	s.ServeJSON()
}

// 添加学生
// @router /:schoolName/students [post]
func (s *SchoolController) StudentPost() {
	schoolName := s.GetString(":schoolName")
	fmt.Printf(schoolName)
	var student models.Student
	json.Unmarshal(s.Ctx.Input.RequestBody, &student)
	studentName := models.AddStudent(student)
	associate := models.Associate{schoolName, student.Name}
	models.AddAssociate(associate)
	s.Data["json"] = map[string]string{"studentName": studentName}
	s.ServeJSON()
}

// 查询关系
// @router /associate [get]
func (s *SchoolController) Associate() {
	associate := models.GetAllAssociate()
	s.Data["json"] = associate
	s.ServeJSON()
}