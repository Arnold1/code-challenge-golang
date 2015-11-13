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
	var students []string
	for _, s := range f.students {
		students = append(students, fmt.Sprintf("%v %v %v %v %v", s.LastName(), s.FirstName(), s.Campus(), s.DateOfBirth(), s.FavoriteColor()))
	}
	return strings.Join(students, "\n")
}
