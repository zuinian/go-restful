package controllers

import (
	"github.com/astaxie/beego"
	"srm/srm-apiserver/models"
	"encoding/json"
)

// Operations about object
type SchoolController struct {
	beego.Controller
}

// 注册学校
// @router / [post]
func (s *SchoolController) Post() {
	var school models.School
	json.Unmarshal(s.Ctx.Input.RequestBody, &school)
	sName := models.AddSchool(school)
	s.Data["json"] = map[string]string{"sName": sName}
	s.ServeJSON()
}

// 查询所有学校
// @router / [get]
func (s *SchoolController) GetAll() {
	schools := models.GetAllSchools()
	s.Data["json"] = schools
	s.ServeJSON()
}

// 查询某个学校所有学生
// @router /:schoolName/students [get]
func (s *SchoolController) StudentsBySchoolGet() {
	schoolName := s.GetString(":schoolName")
	students := models.AssociateBySchoolGet(schoolName)
	s.Data["json"] = students
	s.ServeJSON()
}

// 添加学生
// @router /:schoolName/students [post]
func (s *SchoolController) StudentPost() {
	schoolName := s.GetString(":schoolName")
	// 添加学生信息
	var student models.Student
	json.Unmarshal(s.Ctx.Input.RequestBody, &student)
	studentName := models.StudentAdd(student)

	// 所在学校绑定该学生
	associate := models.Associate{schoolName, student.Name}
	models.AssociateAdd(associate)

	s.Data["json"] = map[string]string{"schoolName": schoolName,"studentName": studentName}
	s.ServeJSON()
}

// 变更学生信息
// @router /:schoolName/students/:studentName [put]
func (s *SchoolController) StudentPut() {
	schoolName := s.GetString(":schoolName")
	studentName := s.GetString(":studentName")
	if !models.AssociateExist(schoolName, studentName) {
		s.Data["json"] = "该学生不存在"
		s.ServeJSON()
		return
	}

	var student models.Student
	json.Unmarshal(s.Ctx.Input.RequestBody, &student)
	bool := models.StudentPut(student)
	if bool {
		s.Data["json"] = "修改成功"
	} else {
		s.Data["json"] = "修改失败，该生不存在"
	}
	s.ServeJSON()
}

// 查询指定学生
// @router /:schoolName/students/:studentName [Get]
func (s *SchoolController) StudentGet() {
	schoolName := s.GetString(":schoolName")
	studentName := s.GetString(":studentName")
	if !models.AssociateExist(schoolName, studentName) {
		s.Data["json"] = "该学生不存在该校"
		s.ServeJSON()
		return
	}

	student := models.StudentGet(studentName)
	s.Data["json"] = student
	s.ServeJSON()
}

// 删除指定学生
// @router /:schoolName/students/:studentName [Delete]
func (s *SchoolController) StudentDelete() {
	schoolName := s.GetString(":schoolName")
	studentName := s.GetString(":studentName")
	if !models.AssociateExist(schoolName, studentName) {
		s.Data["json"] = "该学生不存在该校"
		s.ServeJSON()
		return
	}

	bool := models.StudentDelete(studentName)
	if !bool {
		s.Data["json"] = "无该学生信息"
		s.ServeJSON()
		return
	}

	bool2 := models.AssociateDelete(schoolName, studentName)
	if !bool2 {
		s.Data["json"] = "删除失败，可能该生不属于该校"
		s.ServeJSON()
		return
	}
	s.Data["json"] = "删除成功"
	s.ServeJSON()
}

// 学生转学
// @router /:schoolNameOld/students/:studentName/:schoolNameNew [Put]
func (s *SchoolController) StudentTransfer() {
	schoolNameOld := s.GetString(":schoolNameOld")
	schoolNameNew := s.GetString(":schoolNameNew")
	studentName := s.GetString(":studentName")
	bool := models.AssociateTranser(schoolNameOld, schoolNameNew, studentName)
	if !bool {
		s.Data["json"] = "转学失败"
		s.ServeJSON()
		return
	}
	s.Data["json"] = "转学成功"
	s.ServeJSON()
}

// 查询关系
// @router /associate [get]
func (s *SchoolController) Associate() {
	associate := models.AssociateAllGet()
	s.Data["json"] = associate
	s.ServeJSON()
}