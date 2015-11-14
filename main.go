package main

import (
	"os"

	"github.com/arnoldcano/code-challenge2/students"
)

func main() {
	options := students.OptionGroup{
		"output1": students.Option{
			"file":      "data/comma.txt",
			"delimiter": students.CommaDelimiter,
		},
		"output2": students.Option{
			"file":      "data/dollar.txt",
			"delimiter": students.DollarDelimiter,
		},
		"output3": students.Option{
			"file":      "data/pipe.txt",
			"delimiter": students.PipeDelimiter,
		},
	}
	students.NewRunner(os.Stdout, options).Run()
}
