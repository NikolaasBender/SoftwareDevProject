package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// var templates = template.Must(template.ParseGlob("templates/*"))

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	//will eventually make random key generation
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

//THIS HANDLES ANYTHING IN THE TOP DIRECTORY
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	pathVariables := mux.Vars(r)
	fmt.Println("HOME HANDLER: '" + pathVariables["page"] + "'" + "'" + r.URL.Path + "'")
	PageVars := PageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, _ := template.ParseFiles(pathVariables["page"])

	t.Execute(w, PageVars)
}

//SUPER BASIC INDEX HANDLER
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "/view/index.html")
}

//THIS SHOULD HANDLE THE LOGIN
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("login_test.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	session, _ := store.Get(r, "login_cookie")
	// Authentication goes here
	// ...

	details := ContactDetails{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// Set user as authenticated
	session.Values["username"] = details.Username
	session.Values["loggedIn"] = true
	session.Save(r, w)

	// do something with details
	fmt.Println(session.Values["username"])

	t.Execute(w, struct{ Success bool }{true})
}

var titles = []string{"t1", "t2", "t3", "t4"}
var contents = []string{"c1", "c2", "c3", "c4"}

//OUR ATTEMPT FOR DEALING WITH CARDS
func CardHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	session, _ := store.Get(r, "login_cookie")
	fmt.Println(session.Values["username"])
	p := Card{Title: "", Content: ""}
	if session.Values["loggedIn"] != true {
		p = Card{Title: "you're not logged in", Content: "Please login"}
	} else {
		p = Card{Title: "you are logged in", Content: "Your stuff is due soon"}
	}
	t, _ := template.ParseFiles("templates/newcard.html")
	t.Execute(w, p)
}

//DEMO FOR DEALING WITH FORMS
func FormHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("formTest.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	details := BigForm{
		nickname:  r.FormValue("nickname"),
		email:     r.FormValue("email"),
		password:  r.FormValue("password"),
		gender:    r.FormValue("gender"),
		securityQ: r.FormValue("securityQuestion"),
		languages: r.FormValue("Languages"),
		textbox:   r.FormValue("textbox"),
	}

	fmt.Println(details)

	t.Execute(w, struct{ Success bool }{true})

}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	fmt.Println("VIEW HANDLER: '" + pathVariables["page"] + "'" + "'" + r.URL.Path + "'")

	page := ""

	if strings.Contains(pathVariables["page"], ".html") == true {
		page = pathVariables["page"]
	} else {
		page = pathVariables["page"] + ".html"
	}

	//READ CSS FILE
	//SAVE ALL OF THAT TO A VAIRAIBLE
	//SAVE IT IN p
	//MODIFTY HTML A LITTLE BIT

	p := Page{footer: readStyle}
	err := views.ExecuteTemplate(w, page, p)
	if err != nil {
		log.Fatal("Cannot Get View ", err)
	}
}

func readStyle string (){
	b, err := ioutil.ReadFile("assets/style.css") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'
}

// func AssetsHandler(w http.ResponseWriter, r *http.Request) {
// 	now := time.Now()
// 	pathVariables := mux.Vars(r)
// 	fmt.Println("ASSETS HANDLER: '" + pathVariables["page"] + "'" + "'" + r.URL.Path + "'")
// 	PageVars := PageVariables{ //store the date and time in a struct
// 		Date: now.Format("02-01-2006"),
// 		Time: now.Format("15:04:05"),
// 	}

// 	t, err := template.ParseFiles("assets/" + pathVariables["page"] + ".css")

// 	fmt.Println(err.Error())

// 	t.Execute(w, PageVars)
// }
