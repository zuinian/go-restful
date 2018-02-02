package controllers

import (
	"net/http"
	"github.com/astaxie/beego"
	"srm/srm-apiserver/models"
	"encoding/json"
)

// Operations about object
type StudentController struct {
	beego.Controller
}

// 添加学生
func (s *StudentController) Post() {
	// 获取学生信息
	var student models.Student
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &student)
	if err != nil {
		s.Data["json"] = ResponseError{"Invaild parameter param", 400,}
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		s.ServeJSON()
		return
	}
	// 获取学籍信息
	var schoolName struct {SchoolName string `json:"schoolName"`}
	json.Unmarshal(s.Ctx.Input.RequestBody, &schoolName)

	// 添加学生
	models.StudentAdd(student)

	// 添加学籍
	associate := models.Associate{schoolName.SchoolName, student.Name}
	models.AssociateAdd(associate)
}

// 变更学生信息
func (s *StudentController) Put() {
	// studentName := s.GetString(":studentName")
	// 获取学生信息
	var student models.Student
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &student)
	if err != nil {
		s.Data["json"] = ResponseError{"Invaild parameter param", 400,}
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		s.ServeJSON()
		return
	}

	// 修改学生信息
	ok := models.StudentPut(student)
	if !ok {
		s.Data["json"] = ResponseError{"The student dost's exist", 400,}
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		s.ServeJSON()
	}
}

// 查询指定学生
func (s *StudentController) Get() {
	studentName := s.GetString(":studentName")
	student := models.StudentGet(studentName)
	if student == nil {
		s.Data["json"] = ResponseError{"The student dost's exist", 400,}
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		s.ServeJSON()
		return
	}
	s.Data["json"] = student
	s.ServeJSON()
}

// 删除某个学生
func (s *StudentController) Delete() {
	studentName := s.GetString(":studentName")

	// 删除学生信息
	ok := models.StudentDelete(studentName)
	if !ok {
		s.Data["json"] = ResponseError{"The student dost's exist", 400,}
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		s.ServeJSON()
		return
	}

	// 删除学籍
	ok2 := models.AssociateDelete(studentName)
	if !ok2 {
		s.Data["json"] = ResponseError{"学籍信息删除失败，未查询到学籍信息", 400,}
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		s.ServeJSON()
		return
	}
}
