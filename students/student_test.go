package students

import "testing"

func TestNewStudent(t *testing.T) {
	data := map[string]string{
		"firstName":     "a",
		"middleInitial": "b",
		"lastName":      "c",
		"campus":        "d",
		"favoriteColor": "e",
		"dateOfBirth":   "1/1/2015",
	}
	s := NewStudent(data)

	if s.FirstName() != "a" {
		t.Errorf("Expected '%v', got '%v'", "a", s.FirstName())
	}
	if s.MiddleInitial() != "b" {
		t.Errorf("Expected '%v', got '%v'", "b", s.MiddleInitial())
	}
	if s.LastName() != "c" {
		t.Errorf("Expected '%v', got '%v'", "c", s.LastName())
	}
	if s.Campus() != "d" {
		t.Errorf("Expected '%v', got '%v'", "d", s.Campus())
	}
	if s.FavoriteColor() != "e" {
		t.Errorf("Expected '%v', got '%v'", "e", s.FavoriteColor())
	}
	if s.DateOfBirth() != "1/1/2015" {
		t.Errorf("Expected '%v', got '%v'", "1/1/2015", s.DateOfBirth())
	}
}

func TestTrimsExtraSpaces(t *testing.T) {
	data := map[string]string{
		"firstName":     " a ",
		"middleInitial": " b ",
		"lastName":      " c ",
		"campus":        " d ",
		"favoriteColor": " e ",
		"dateOfBirth":   " 1/1/2015 ",
	}
	s := NewStudent(data)

	if s.FirstName() != "a" {
		t.Errorf("Expected '%v', got '%v'", "a", s.FirstName())
	}
	if s.MiddleInitial() != "b" {
		t.Errorf("Expected '%v', got '%v'", "b", s.MiddleInitial())
	}
	if s.LastName() != "c" {
		t.Errorf("Expected '%v', got '%v'", "c", s.LastName())
	}
	if s.Campus() != "d" {
		t.Errorf("Expected '%v', got '%v'", "d", s.Campus())
	}
	if s.FavoriteColor() != "e" {
		t.Errorf("Expected '%v', got '%v'", "e", s.FavoriteColor())
	}
	if s.DateOfBirth() != "1/1/2015" {
		t.Errorf("Expected '%v', got '%v'", "1/1/2015", s.DateOfBirth())
	}
}

func TestBadCampuses(t *testing.T) {
	s := new(Student)
	s.SetCampus("LA")

	if s.Campus() != "Los Angeles" {
		t.Errorf("Expected '%v', got '%v'", "Los Angeles", s.Campus())
	}
	s.SetCampus("NYC")
	if s.Campus() != "New York City" {
		t.Errorf("Expected '%v', got '%v'", "New York City", s.Campus())
	}
	s.SetCampus("SF")
	if s.Campus() != "San Francisco" {
		t.Errorf("Expected '%v', got '%v'", "San Francisco", s.Campus())
	}
}

func TestBadDateOfBirth(t *testing.T) {
	s := new(Student)
	s.SetDateOfBirth("1-1-2015")

	if s.DateOfBirth() != "1/1/2015" {
		t.Errorf("Expected '%v', got '%v'", "1/1/2015", s.DateOfBirth())
	}
}
