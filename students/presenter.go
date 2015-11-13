package students

import (
	"sort"
	"time"
)

const dateOfBirthFormat = "1/2/2006"

type sorter func(s1, s2 *Student) bool

type Presenter struct {
	students []*Student
	sorters  []sorter
}

func NewPresenter(students []*Student) *Presenter {
	return &Presenter{
		students: students,
	}
}

func (p *Presenter) ByCampusAndLastNameAsc() []*Student {
	campus := func(s1, s2 *Student) bool {
		return s1.Campus() < s2.Campus()
	}
	lastName := func(s1, s2 *Student) bool {
		return s1.LastName() < s2.LastName()
	}
	sortBy(campus, lastName).Sort(p.students)
	return p.students
}

func (p *Presenter) ByDateOfBirthAsc() []*Student {
	dateOfBirth := func(s1, s2 *Student) bool {
		d1, _ := time.Parse(dateOfBirthFormat, s1.DateOfBirth())
		d2, _ := time.Parse(dateOfBirthFormat, s2.DateOfBirth())
		return d1.Before(d2)
	}
	sortBy(dateOfBirth).Sort(p.students)
	return p.students
}

func (p *Presenter) ByLastNameDesc() []*Student {
	lastName := func(s1, s2 *Student) bool {
		return s1.LastName() > s2.LastName()
	}
	sortBy(lastName).Sort(p.students)
	return p.students
}

func (p *Presenter) Sort(students []*Student) {
	p.students = students
	sort.Sort(p)
}

func (p *Presenter) Len() int {
	return len(p.students)
}

func (p *Presenter) Swap(i, j int) {
	p.students[i], p.students[j] = p.students[j], p.students[i]
}

func (p *Presenter) Less(i, j int) bool {
	var k int

	s1 := p.students[i]
	s2 := p.students[j]
	for k = 0; k < len(p.sorters)-1; k++ {
		sorter := p.sorters[k]
		switch {
		case sorter(s1, s2):
			return true
		case sorter(s2, s1):
			return false
		}
	}
	return p.sorters[k](s1, s2)
}

func sortBy(sorters ...sorter) *Presenter {
	return &Presenter{
		sorters: sorters,
	}
}
