package models

var (
	associateList map[string]map[string]bool
)

func init() {
	associateList = make(map[string]map[string]bool)
	a := Associate{"xiada","xiaoming"}
	associateList[a.SchoolName] = make(map[string]bool)
	associateList[a.SchoolName][a.StudentName] = true
}

type Associate struct {
	SchoolName string
	StudentName string
}

func AddAssociate(a Associate) {
	if (associateList[a.SchoolName] == nil) {
		associateList[a.SchoolName] = make(map[string]bool)
	}
	associateList[a.SchoolName][a.StudentName] = true
}

func GetAllAssociate() map[string]map[string]bool {
	return associateList
}
