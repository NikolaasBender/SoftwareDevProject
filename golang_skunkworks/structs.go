package main

//"fmt"
//"html/template"
//"log"
//"net/http"
//"time"

//"github.com/gorilla/mux"

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


//WE WILL WANT TO GET MORE STUFF ABOUT EACH POST
type Post struct {
	Title   string
	Content string
	Link    string
	Type	int
}

type Page struct {
	style string
}
