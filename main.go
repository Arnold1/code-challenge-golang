package main

import (
	"os"

	"github.com/arnoldcano/code-challenge2/students"
)

func main() {
	options := map[string]string{
		"commaFile":       "data/comma.txt",
		"dollarFile":      "data/dollar.txt",
		"pipeFile":        "data/pipe.txt",
		"commaDelimiter":  ",",
		"dollarDelimiter": "$",
		"pipeDelimiter":   "|",
	}
	students.NewRunner(os.Stdout, options).Run()
}
