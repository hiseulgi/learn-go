package main

// menggunakan pointer karena hanya memakai struct Student sekali saja
var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

// fungsi init() -> fungsi yang secara otomatis dipanggil ketika package tersebut di import atau di run
func init() {
	students = append(students, &Student{"s001", "Sugab", 2})
	students = append(students, &Student{"s002", "Riski", 2})
	students = append(students, &Student{"s003", "Anam", 2})
}
