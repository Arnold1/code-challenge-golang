package students

import (
	"fmt"
	"io"
	"os"
)

type Option map[string]string
type Options []Option

type Runner struct {
	out     io.Writer
	options Options
}

func NewRunner(out io.Writer, options Options) *Runner {
	return &Runner{
		out:     out,
		options: options,
	}
}

func (r *Runner) Run() {
	r.display(r.students())
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

func (r *Runner) students() []*Student {
	var s []*Student

	for _, o := range r.options {
		f, _ := os.Open(o["file"])
		defer f.Close()
		r := NewReader(f, o["delimiter"])
		s = append(s, r.Read()...)
	}
	return s
}
