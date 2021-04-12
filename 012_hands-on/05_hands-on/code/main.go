package main

import (
	"log"
	"os"
	"text/template"
)

type meal struct {
	MealType string
	Items    []string
}

type restaurant struct {
	Name string
	Menu []meal
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	r := restaurant{
		Name: "Gopal's",
		Menu: []meal{
			{
				MealType: "Breakfast",
				Items: []string{
					"Poha",
					"Dosa",
					"Idly",
				},
			},
			{
				MealType: "Lunch",
				Items: []string{
					"Curd Rice",
					"Veg Thali",
					"Dosa",
				},
			},
			{
				MealType: "Dinner",
				Items: []string{
					"Bissebeli bath",
					"Tiramisu",
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", r)
	if err != nil {
		log.Fatalln(err)
	}
}
