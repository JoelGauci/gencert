package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/gencert/cert"
	"training.go/gencert/html"
	"training.go/gencert/pdf"
)

func main() {
	file := flag.String("file", "", "CSV input file")
	outputType := flag.String("type", "pdf", "Output type.")
	flag.Parse()

	if len(*file) <= 0 {
		fmt.Printf("Invalid file. got='%v'\n", *file)
		os.Exit(1)
	}

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown output type. got=%v\n", *outputType)
	}

	if err != nil {
		fmt.Printf("Could not create pdf saver. got=%v\n", err)
		os.Exit(1)
	}

	certs, err := cert.ParseCSV(*file)
	if err != nil {
		fmt.Printf("Could not parse CSV file. got=%v\n", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save Cert. got=%v\n", err)
		}
	}

}
