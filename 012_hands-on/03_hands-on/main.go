package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{
			Name:    "Hotel One",
			Address: "123 Way",
			City:    "Stockton",
			Zip:     95209,
			Region:  "Central",
		},
		hotel{
			Name:    "Hotel Two",
			Address: "345 st",
			City:    "Lodi",
			Zip:     95241,
			Region:  "Central",
		},
	}

	err := tmpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
