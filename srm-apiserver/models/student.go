package models

var (
	StudentList map[string]*Student
)

func init() {
	StudentList = make(map[string]*Student)
	s := Student{"xiaoming","one", "one"}
	StudentList["0"] = &s
}

type Student struct {
	Name string
	Grade string
	Class string
}

func AddStudent(s Student) string {
	StudentList[s.Name] = &s
	return s.Name
}

func GetAllStudents() map[string]*Student {
	return StudentList
}