package controllers

import (
	"net/http"
	"github.com/astaxie/beego"
	"srm/srm-apiserver/models"
	"encoding/json"
)

// Operations about object
type SchoolController struct {
	beego.Controller
}

type ResponseError struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

// 注册学校
func (s *SchoolController) Post() {
	var school models.School
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &school)
	if err != nil {
		s.Data["json"] = ResponseError{"Invaild parameter param", 400,}
		s.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		s.ServeJSON()
		return
	}
	models.SchoolPost(school)
}

// 查询所有学校
func (s *SchoolController) GetAll() {
	schoolsMap := models.SchoolsGet()
	schools := make([]*models.School, 0)
	for k,_ := range schoolsMap {
		schools = append(schools, schoolsMap[k])
	}
	s.Data["json"] = schools
	s.ServeJSON()
}

// 查询某个学校所有学生
func (s *SchoolController) StudentsBySchoolGet() {
	schoolName := s.GetString(":schoolName")
	associateMap := models.AssociateBySchoolGet(schoolName)
	studentMap := models.StudentsAllGet()
	// 遍历学生姓名获取详细信息
	students := make([]models.Student, 0)
	for studentName, _ := range associateMap {
		students = append(students, *studentMap[studentName])
	}
	s.Data["json"] = students
	s.ServeJSON()
}
