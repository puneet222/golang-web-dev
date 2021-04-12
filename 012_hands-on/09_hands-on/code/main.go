package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	f, err := os.Open("../table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	var headers []string = lines[0]
	var data [][]string = lines[1:]
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", struct {
		Headers []string
		Data [][]string
	}{
		Headers: headers,
		Data: data,
	})
}
