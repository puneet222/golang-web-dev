package main

import (
	"log"
	"os"
	"text/template"
)

type Region int

const (
	Northern Region = iota
	Southern
	Central
	)

func (r Region) String() string {
	return [...]string{"Northern", "Southern", "Central"}[r]
}

type hotel struct {
	Name string
	Address string
	City string
	Zip int
	Region Region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	hotels := []hotel{
		{
			Name: "Hotel 1",
			Address: "123, street 1",
			City: "california",
			Zip: 3223,
			Region: Central,
		},
		{
			Name: "Hotel 2",
			Address: "456, street 2",
			City: "San Fransisco",
			Zip: 6767,
			Region: Northern,
		},
		{
			Name: "Hotel 3",
			Address: "099, street 00",
			City: "Honolulu",
			Zip: 9090,
			Region: Southern,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
