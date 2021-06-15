package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
	w.Write([]byte("test"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":8099", h)
}
