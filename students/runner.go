package students

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

const ()

type Runner struct {
	out     io.Writer
	options map[string]string
}

func NewRunner(out io.Writer, options map[string]string) *Runner {
	return &Runner{
		out:     out,
		options: options,
	}
}

func (r *Runner) Run() {
	r.display(r.allStudents())
}

func (r *Runner) display(s []*Student) {
	p := NewPresenter(s)

	fmt.Fprintln(r.out, "Output 1:")
	fmt.Fprintln(r.out, NewFormatter(p.ByCampusAndLastNameAsc()).Format())
	fmt.Fprintln(r.out)

	fmt.Fprintln(r.out, "Output 2:")
	fmt.Fprintln(r.out, NewFormatter(p.ByDateOfBirthAsc()).Format())
	fmt.Fprintln(r.out)

	fmt.Fprintln(r.out, "Output 3:")
	fmt.Fprintln(r.out, NewFormatter(p.ByLastNameDesc()).Format())
}

func (r *Runner) allStudents() []*Student {
	cf, _ := os.Open(r.options["commaFile"])
	df, _ := os.Open(r.options["dollarFile"])
	pf, _ := os.Open(r.options["pipeFile"])
	defer cf.Close()
	defer df.Close()
	defer pf.Close()

	cd, _ := utf8.DecodeRuneInString(r.options["commaDelimiter"])
	dd, _ := utf8.DecodeRuneInString(r.options["dollarDelimiter"])
	pd, _ := utf8.DecodeRuneInString(r.options["pipeDelimiter"])

	cr := NewReader(cf, cd)
	dr := NewReader(df, dd)
	pr := NewReader(pf, pd)

	var all []*Student
	all = append(all, cr.Read()...)
	all = append(all, dr.Read()...)
	all = append(all, pr.Read()...)

	return all
}
