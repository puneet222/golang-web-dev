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

func getCsvRecords(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	return records
}

func main() {
	records := getCsvRecords("../table.csv")
	var headers = records[0]
	var data = records[1:]
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", struct {
		Headers []string
		Data [][]string
	}{
		Headers: headers,
		Data: data,
	})
}
