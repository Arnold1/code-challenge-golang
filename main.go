package main

import (
	"os"

	"github.com/arnoldcano/code-challenge2/students"
)

func main() {
	options := students.Options{
		students.Option{
			"file":      "data/comma.txt",
			"delimiter": ",",
		},
		students.Option{
			"file":      "data/dollar.txt",
			"delimiter": "$",
		},
		students.Option{
			"file":      "data/pipe.txt",
			"delimiter": "|",
		},
	}
	students.NewRunner(os.Stdout, options).Run()
}
