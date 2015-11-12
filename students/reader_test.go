package students

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	in := `"a","b","c","d","e","1/1/2015"`
	delimiter := ','
	r := NewReader(strings.NewReader(in), delimiter)

	if r.Delimiter != delimiter {
		t.Error("Expected '%v', got '%v'", delimiter, r.Delimiter)
	}
}

func TestCommaRead(t *testing.T) {
	data := map[string]string{
		"firstName":     "a",
		"lastName":      "c",
		"campus":        "d",
		"favoriteColor": "e",
		"dateOfBirth":   "1/1/2015",
	}
	s := NewStudent(data)
	expected := []*Student{s}

	in := `"c","a","d","e","1/1/2015"`
	delimiter := ','
	r := NewReader(strings.NewReader(in), delimiter)
	students := r.Read()

	if !reflect.DeepEqual(students, expected) {
		t.Errorf("Expected '%v', got '%v'", expected, students)
	}
}

func TestDollarRead(t *testing.T) {
	data := map[string]string{
		"firstName":     "a",
		"middleInitial": "b",
		"lastName":      "c",
		"campus":        "d",
		"favoriteColor": "e",
		"dateOfBirth":   "1/1/2015",
	}
	s := NewStudent(data)
	expected := []*Student{s}

	in := `"c"$"a"$"b"$"d"$"1/1/2015"$"e"`
	delimiter := '$'
	r := NewReader(strings.NewReader(in), delimiter)
	students := r.Read()

	if !reflect.DeepEqual(students, expected) {
		t.Errorf("Expected '%v', got '%v'", expected, students)
	}
}

func TestPipeRead(t *testing.T) {
	data := map[string]string{
		"firstName":     "a",
		"middleInitial": "b",
		"lastName":      "c",
		"campus":        "d",
		"favoriteColor": "e",
		"dateOfBirth":   "1/1/2015",
	}
	s := NewStudent(data)
	expected := []*Student{s}

	in := `"c"|"a"|"b"|"d"|"e"|"1/1/2015"`
	delimiter := '|'
	r := NewReader(strings.NewReader(in), delimiter)
	students := r.Read()

	if !reflect.DeepEqual(students, expected) {
		t.Errorf("Expected '%v', got '%v'", expected, students)
	}
}
