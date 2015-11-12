package main

import (
	"fmt"
	"os"

	"github.com/arnoldcano/code-challenge2/students"
)

const (
	commaFile  = "data/comma.txt"
	dollarFile = "data/dollar.txt"
	pipeFile   = "data/pipe.txt"

	commaDelimiter  = ','
	dollarDelimiter = '$'
	pipeDelimiter   = '|'
)

func main() {
	display(output())
}

func display(s []*students.Student) {
	p := students.NewPresenter(s)

	fmt.Println("Output 1:")
	fmt.Println(students.NewFormatter(p.ByCampusAndLastNameAsc()).Format())
	fmt.Println()

	fmt.Println("Output 2:")
	fmt.Println(students.NewFormatter(p.ByDateOfBirthAsc()).Format())
	fmt.Println()

	fmt.Println("Output 3:")
	fmt.Println(students.NewFormatter(p.ByLastNameDesc()).Format())
}

func output() []*students.Student {
	cf, _ := os.Open(commaFile)
	df, _ := os.Open(dollarFile)
	pf, _ := os.Open(pipeFile)
	defer cf.Close()
	defer df.Close()
	defer pf.Close()

	cr := students.NewReader(cf, commaDelimiter)
	dr := students.NewReader(df, dollarDelimiter)
	pr := students.NewReader(pf, pipeDelimiter)

	var s []*students.Student
	s = append(s, cr.Read()...)
	s = append(s, dr.Read()...)
	s = append(s, pr.Read()...)

	return s
}
