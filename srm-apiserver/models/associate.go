package models

import "fmt"

var (
	associateList map[string]map[string]bool
)

func init() {
	associateList = make(map[string]map[string]bool)
	a := Associate{"xiada","xiaoming"}
	AssociateAdd(a)

	associateList["fuda"] = make(map[string]bool)
}

type Associate struct {
	SchoolName string `json:"schoolName"`
	StudentName string `json:"studentName"`
}

// 学籍添加
func AssociateAdd(a Associate) bool {
	if (associateList[a.SchoolName] == nil) {
		associateList[a.SchoolName] = make(map[string]bool)
	}
	associateList[a.SchoolName][a.StudentName] = true
	return true
}

// 学籍查询--全部
func AssociateAllGet() map[string]map[string]bool {
	return associateList
}

// 学籍查询-某个学校
func AssociateBySchoolGet(schoolName string) map[string]bool {
	return associateList[schoolName]
}
// 学籍查询判断-某个学校是否有某个学生
func AssociateExist(schoolName string, studentName string) bool {
	if associateList[schoolName] != nil && associateList[schoolName][studentName] {
		return true
	}
	return false
}

// 学籍删除
func AssociateDelete(studentName string) bool {
	for schoolName, _ := range associateList {
		if associateList[schoolName][studentName] == true {
			delete(associateList[schoolName], studentName)
			return true
		}
	}

	return false
}

// 学籍修改-转学
func AssociatePut(a Associate) bool {
	// 先获取旧学校
	schoolNameOld := ""
	for schoolName, _ := range associateList {
		if associateList[schoolName][a.StudentName] == true {
			schoolNameOld = schoolName
			break
		}
	}
	fmt.Print("ddd")
	fmt.Print(schoolNameOld)
	if (schoolNameOld == "" || schoolNameOld == a.SchoolName) {
		return false
	}

	// 变更学籍信息-删除旧的再添加新的就好了
	delete(associateList[schoolNameOld], a.StudentName)
	return AssociateAdd(a)
}
