package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./starting-files/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("./starting-files/public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/dog", dog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
}