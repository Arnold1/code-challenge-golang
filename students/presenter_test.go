package students

import (
	"reflect"
	"testing"
)

func TestByCampusAndLastNameAsc(t *testing.T) {
	data1 := map[string]string{
		"firstName":     "",
		"middleInitial": "",
		"lastName":      "b",
		"campus":        "d",
		"favoriteColor": "",
		"dateOfBirth":   "",
	}
	data2 := map[string]string{
		"firstName":     "",
		"middleInitial": "",
		"lastName":      "a",
		"campus":        "d",
		"favoriteColor": "",
		"dateOfBirth":   "",
	}
	data3 := map[string]string{
		"firstName":     "",
		"middleInitial": "",
		"lastName":      "c",
		"campus":        "a",
		"favoriteColor": "",
		"dateOfBirth":   "",
	}
	s1 := NewStudent(data1)
	s2 := NewStudent(data2)
	s3 := NewStudent(data3)
	sortedStudents := []*Student{s3, s2, s1}
	p := NewPresenter([]*Student{s1, s2, s3})
	students := p.ByCampusAndLastNameAsc()
	if !reflect.DeepEqual(students, sortedStudents) {
		t.Errorf("Expected '%v', got '%v'", sortedStudents, students)
	}
}

func TestByDateOfBirthAsc(t *testing.T) {
	data1 := map[string]string{
		"firstName":     "",
		"middleInitial": "",
		"lastName":      "",
		"campus":        "",
		"favoriteColor": "",
		"dateOfBirth":   "1/1/2015",
	}
	data2 := map[string]string{
		"firstName":     "",
		"middleInitial": "",
		"lastName":      "",
		"campus":        "",
		"favoriteColor": "",
		"dateOfBirth":   "2/1/2015",
	}
	s1 := NewStudent(data1)
	s2 := NewStudent(data2)
	sortedStudents := []*Student{s1, s2}
	p := NewPresenter([]*Student{s2, s1})
	students := p.ByDateOfBirthAsc()
	if !reflect.DeepEqual(students, sortedStudents) {
		t.Errorf("Expected '%v', got '%v'", sortedStudents, students)
	}
}

func TestByLastNameDesc(t *testing.T) {
	data1 := map[string]string{
		"firstName":     "",
		"middleInitial": "",
		"lastName":      "a",
		"campus":        "",
		"favoriteColor": "",
		"dateOfBirth":   "",
	}
	data2 := map[string]string{
		"firstName":     "",
		"middleInitial": "",
		"lastName":      "b",
		"campus":        "",
		"favoriteColor": "",
		"dateOfBirth":   "",
	}
	s1 := NewStudent(data1)
	s2 := NewStudent(data2)
	sortedStudents := []*Student{s2, s1}
	p := NewPresenter([]*Student{s1, s2})
	students := p.ByLastNameDesc()
	if !reflect.DeepEqual(students, sortedStudents) {
		t.Errorf("Expected '%v', got '%v'", sortedStudents, students)
	}
}
