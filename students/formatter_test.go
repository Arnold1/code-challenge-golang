package students

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewFormatter(t *testing.T) {
	data := map[string]string{
		"firstName":     "a",
		"lastName":      "c",
		"campus":        "d",
		"favoriteColor": "e",
		"dateOfBirth":   "1/1/2015",
	}
	s := NewStudent(data)
	students := []*Student{s}

	f := NewFormatter(students)

	if !reflect.DeepEqual(f.students, students) {
		t.Errorf("Expected '%v', got '%v'", students, f.students)
	}
}

func TestFormat(t *testing.T) {
	data1 := map[string]string{
		"firstName":     "a",
		"lastName":      "c",
		"campus":        "d",
		"favoriteColor": "e",
		"dateOfBirth":   "1/1/2015",
	}
	data2 := map[string]string{
		"firstName":     "e",
		"lastName":      "d",
		"campus":        "c",
		"favoriteColor": "a",
		"dateOfBirth":   "2/1/2015",
	}
	s1 := NewStudent(data1)
	s2 := NewStudent(data2)
	students := []*Student{s1, s2}
	expected := strings.Join([]string{
		"c a d 1/1/2015 e",
		"d e c 2/1/2015 a",
	}, "\n")

	f := NewFormatter(students)
	formatted := f.Format()

	if formatted != expected {
		t.Errorf("Expected '%v', got '%v'", expected, formatted)
	}
}
