package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type tableRow struct {
	Date   string
	Values values
}

type values struct {
	Open   string
	High   string
	Low    string
	Close  string
	Volume string
	Adj    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	file, err := ioutil.ReadFile("table.csv")
	if err != nil {
		log.Fatalln(err)
	}

	data := string(file)
	dataRows := strings.Split(data, "\n")

	var parsedRows [][]string

	for _, row := range dataRows {
		splitRow := strings.Split(row, ",")
		parsedRows = append(parsedRows, splitRow)

	}

	parsedRows = parsedRows[:len(parsedRows)-1]

	var tableRows []tableRow

	for _, row := range parsedRows {
		tableRows = append(tableRows, tableRow{
			Date: row[0],
			Values: values{
				Open:   row[1],
				High:   row[2],
				Low:    row[3],
				Close:  row[4],
				Volume: row[5],
				Adj:    row[6],
			},
		})
	}

	tpl.Execute(os.Stdout, tableRows)
}
