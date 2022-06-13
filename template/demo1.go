package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type user struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err: ", err)
		return
	}

	user := user{
		Name:   "Alice",
		Gender: "女",
		Age:    20,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err: ", err)
		return
	}
}
