package students

import (
	"fmt"
	"strings"
)

type Formatter struct {
	students []*Student
}

func NewFormatter(students []*Student) *Formatter {
	return &Formatter{
		students: students,
	}
}

func (f *Formatter) Format() string {
	students := make([]string, len(f.students))
	for i, s := range f.students {
		students[i] = fmt.Sprintf("%v %v %v %v %v", s.LastName(), s.FirstName(), s.Campus(), s.DateOfBirth(), s.FavoriteColor())
	}
	return strings.Join(students, "\n")
}
