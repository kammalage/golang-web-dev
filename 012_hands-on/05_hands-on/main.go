package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	m := menu{
		Breakfast: []string{
			"Bacon and Eggs",
			"Toast",
			"Oatmeal",
		},
		Lunch: []string{
			"Jalapeno Burger with Cheese Skirt",
			"Buffalo Stampede Wings with Garlic Fries",
			"Chicken Noodle Soup",
		},
		Dinner: []string{
			"NY Strip Steak with garlic butter and Mashed Potatoes",
			"Pepperoni Pizza",
			"Chicken Katsu Curry",
		},
	}

	err := tmpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}

}
