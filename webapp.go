package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello, World!")
	}
}

func showPepe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/pepe.gtpl")
		t.Execute(w, nil)
	}
}
func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/pepe", showPepe)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
