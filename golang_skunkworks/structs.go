package main

import "html/template"

//"fmt"
//"html/template"
//"log"
//"net/http"
//"time"

//"github.com/gorilla/mux"

var views = template.Must(template.ParseGlob("view/*"))

type PageVariables struct {
	Date string
	Time string
}

type ContactDetails struct {
	Username string
	Password string
}

type BigForm struct {
	nickname  string
	email     string
	password  string
	gender    string
	securityQ string
	languages string
	textbox   string
}

type Card struct {
	Title   string
	Content string
}

type Page struct {
	footer string
}
