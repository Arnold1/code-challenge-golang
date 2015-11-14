package students

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unicode/utf8"
)

type Option map[string]string
type OptionGroup map[string]Option

type Runner struct {
	out     io.Writer
	options OptionGroup
}

func NewRunner(out io.Writer, options OptionGroup) *Runner {
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
	var all []*Student
	var wg sync.WaitGroup

	for _, o := range r.options {
		wg.Add(1)
		go func(file, delimiter string, wg *sync.WaitGroup) {
			defer wg.Done()
			f, _ := os.Open(file)
			defer f.Close()
			d, _ := utf8.DecodeRuneInString(delimiter)
			r := NewReader(f, d)
			all = append(all, r.Read()...)
		}(o["file"], o["delimiter"], &wg)
	}
	wg.Wait()

	return all
}
