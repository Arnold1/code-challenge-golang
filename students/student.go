package students

import "strings"

type Student struct {
	firstName     string
	middleInitial string
	lastName      string
	campus        string
	favoriteColor string
	dateOfBirth   string
}

func NewStudent(data map[string]string) *Student {
	s := new(Student)
	s.SetFirstName(data["firstName"])
	s.SetMiddleInitial(data["middleInitial"])
	s.SetLastName(data["lastName"])
	s.SetCampus(data["campus"])
	s.SetFavoriteColor(data["favoriteColor"])
	s.SetDateOfBirth(data["dateOfBirth"])
	return s
}

func (s Student) FirstName() string {
	return s.firstName
}

func (s *Student) SetFirstName(firstName string) {
	s.firstName = strings.TrimSpace(firstName)
}

func (s Student) MiddleInitial() string {
	return s.middleInitial
}

func (s *Student) SetMiddleInitial(middleInitial string) {
	s.middleInitial = strings.TrimSpace(middleInitial)
}

func (s Student) LastName() string {
	return s.lastName
}

func (s *Student) SetLastName(lastName string) {
	s.lastName = strings.TrimSpace(lastName)
}

func (s Student) Campus() string {
	return s.campus
}

func (s *Student) SetCampus(campus string) {
	c := strings.TrimSpace(campus)
	switch {
	case strings.Contains(c, "LA"):
		c = strings.Replace(c, "LA", "Los Angeles", -1)
	case strings.Contains(c, "NYC"):
		c = strings.Replace(c, "NYC", "New York City", -1)
	case strings.Contains(c, "SF"):
		c = strings.Replace(c, "SF", "San Francisco", -1)
	}
	s.campus = c
}

func (s Student) FavoriteColor() string {
	return s.favoriteColor
}

func (s *Student) SetFavoriteColor(favoriteColor string) {
	s.favoriteColor = strings.TrimSpace(favoriteColor)
}

func (s Student) DateOfBirth() string {
	return s.dateOfBirth
}

func (s *Student) SetDateOfBirth(dateOfBirth string) {
	dob := strings.TrimSpace(dateOfBirth)
	if strings.Contains(dob, "-") {
		dob = strings.Replace(dob, "-", "/", -1)
	}
	s.dateOfBirth = dob
}
