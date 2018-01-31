package models

var (
	associateList map[string]map[string]bool
)

func init() {
	associateList = make(map[string]map[string]bool)
	a := Associate{"xiada","xiaoming"}
	associateList[a.SchoolName] = make(map[string]bool)
	associateList[a.SchoolName][a.StudentName] = true

	associateList["fuda"] = make(map[string]bool)
}

type Associate struct {
	SchoolName string
	StudentName string
}

// 添加学生关系
func AssociateAdd(a Associate) bool {
	if (associateList[a.SchoolName] == nil) {
		associateList[a.SchoolName] = make(map[string]bool)
	}
	associateList[a.SchoolName][a.StudentName] = true
	return true
}

// 查询所有关系
func AssociateAllGet() map[string]map[string]bool {
	return associateList
}

// 查询某个学校是否有某个学生
func AssociateExist(schoolName string, studentName string) bool {
	if associateList[schoolName] != nil && associateList[schoolName][studentName] {
		return true
	}
	return false
}

// 查询某个学校的所有学生
func AssociateBySchoolGet(schoolName string) map[string]bool {
	return associateList[schoolName]
}

// 删除某个学生
func AssociateDelete(schoolName string, studentName string) bool {
	if associateList[schoolName] != nil && associateList[schoolName][studentName] {
		delete(associateList[schoolName], studentName)
		return true
	}
	return false
}

// 转学
func AssociateTranser(schoolNameOld string, schoolNameNew string, studentName string) bool {
	if associateList[schoolNameOld] != nil && associateList[schoolNameOld][studentName] {
		if associateList[schoolNameNew] != nil {
			delete(associateList[schoolNameOld], studentName)
			a := Associate{schoolNameNew, studentName}
			return AssociateAdd(a)
		}
	}
	return false
}
