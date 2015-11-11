package students

import (
	"encoding/csv"
	"os"
)

type Reader struct {
	file      *os.File
	delimiter rune
}

func NewReader(fileName string, delimiter rune) (*Reader, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := &Reader{
		file:      f,
		delimiter: delimiter,
	}
	return r, nil
}

func (r *Reader) Read() ([]*Student, error) {
	c := csv.NewReader(r.file)
	c.Comma = r.delimiter
	rows, err := c.ReadAll()
	if err != nil {
		return nil, err
	}
	students := make([]*Student, len(rows))
	for _, row := range rows {
		var s *Student
		switch r.delimiter {
		case ',':
			s = newCommaStudent(row)
		case '$':
			s = newDollarStudent(row)
		case '|':
			s = newPipeStudent(row)
		}
		students = append(students, s)
	}
	return students, nil
}

func newCommaStudent(row []string) *Student {
	data := map[string]string{
		"firstName":     row[1],
		"lastName":      row[0],
		"campus":        row[2],
		"favoriteColor": row[3],
		"dateOfBirth":   row[4],
	}
	return NewStudent(data)
}

func newDollarStudent(row []string) *Student {
	data := map[string]string{
		"firstName":     row[1],
		"middleInitial": row[2],
		"lastName":      row[0],
		"campus":        row[3],
		"favoriteColor": row[5],
		"dateOfBirth":   row[4],
	}
	return NewStudent(data)
}

func newPipeStudent(row []string) *Student {
	data := map[string]string{
		"firstName":     row[1],
		"middleInitial": row[2],
		"lastName":      row[0],
		"campus":        row[3],
		"favoriteColor": row[4],
		"dateOfBirth":   row[5],
	}
	return NewStudent(data)
}
