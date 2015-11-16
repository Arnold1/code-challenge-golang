package students

import (
	"encoding/csv"
	"io"
	"unicode/utf8"
)

type Reader struct {
	file      io.Reader
	delimiter rune
}

func NewReader(file io.Reader, delimiter string) *Reader {
	d, _ := utf8.DecodeRuneInString(delimiter)
	return &Reader{
		file:      file,
		delimiter: d,
	}
}

func (r *Reader) Read() []*Student {
	c := csv.NewReader(r.file)
	c.Comma = r.delimiter
	rows, _ := c.ReadAll()
	students := make([]*Student, len(rows))
	for i, row := range rows {
		switch r.delimiter {
		case ',':
			students[i] = newCommaStudent(row)
		case '$':
			students[i] = newDollarStudent(row)
		case '|':
			students[i] = newPipeStudent(row)
		}
	}
	return students
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
