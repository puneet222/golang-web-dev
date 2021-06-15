package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./dog.gohtml"))
}

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/scene/", scene)
	http.HandleFunc("/image.jpg", image)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func scene(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func image(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "image.jpg")
}
