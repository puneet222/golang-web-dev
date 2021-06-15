package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func def(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "def.gohtml", nil)
	HandleError(w, err)
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	HandleError(w, err)
}

func me(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Errorf("error %v", err)
	}
	name := r.URL.Query()
	err = tpl.ExecuteTemplate(w, "me.gohtml", name)
	HandleError(w, err)
}

func main() {
	http.Handle("/", http.HandlerFunc(def))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
