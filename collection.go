package main

import (
	"bytes"

	"github.com/dslipak/pdf"
)

/*
* Collects all the information from the pdf file in order to be validated
 */

func collect(path string) string {
	content, err := readPdf(path) // Read local pdf file
	if err != nil {
		panic(err)
	}
	return content
}

func readPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	// remember close file
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
