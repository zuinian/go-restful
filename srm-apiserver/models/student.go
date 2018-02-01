package models

var (
	StudentList map[string]*Student
)

func init() {
	StudentList = make(map[string]*Student)
	s := Student{"xiaoming","one", 18}
	StudentList["xiaoming"] = &s
}

type Student struct {
	Name   string `json: "name"`
	Gender string `json: "gender"`
	Age    int8 `json: "age"`
}

// 添加学生
func StudentAdd(s Student) string {
	StudentList[s.Name] = &s
	return s.Name
}

// 查询所有学生
func StudentsAllGet() map[string]*Student {
	return StudentList
}

// 查询指定学生
func StudentGet(studentName string) Student {
	return *StudentList[studentName]
}

// 变更学生信息
func StudentPut(s Student) bool {
	if StudentList[s.Name] == nil {
		return false
	}
	StudentList[s.Name] = &s
	return true
}

// 删除学生信息
func StudentDelete(studentName string) bool {
	if StudentList[studentName] == nil {
		return false
	}
	delete(StudentList, studentName)
	return true
}

