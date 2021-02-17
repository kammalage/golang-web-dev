package main

import (
	"html/template"
	"log"
	"os"
)

type restaurant struct {
	Name string
	Menu menu
}

type menu struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	r := []restaurant{
		restaurant{
			Name: "Chilis",
			Menu: menu{
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
			},
		},
		restaurant{
			Name: "Applebees",
			Menu: menu{
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
			},
		},
	}
	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}

}
