package main

import (
	"fmt"
	"html/template"

	//"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

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

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, r.URL.Path[1:])
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("login_test.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	session, _ := store.Get(r, "cookie-name")
	// Authentication goes here
	// ...

	details := ContactDetails{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// Set user as authenticated
	session.Values["username"] = details.Username
	session.Save(r, w)

	// do something with details
	fmt.Println(session.Values["username"])

	t.Execute(w, struct{ Success bool }{true})
}

func CardHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("newCard.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}
	//will add more, just not sure what goes into a card
	details := Card{
		Text: r.FormValue("text"),
	}

	fmt.Println(details)

	t.Execute(w, struct{ Success bool }{true})
}


func FormHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("formTest.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	details := BigForm{
		nickname: r.FormValue("nickname"),
		email: r.FormValue("email"),
		password: r.FormValue("password"),
		gender: r.FormValue("gender"), 
		securityQ: r.FormValue("securityQuestion"), 
		languages: r.FormValue("Languages"),
		textbox: r.FormValue("textbox"),
	}

	fmt.Println(details)

	t.Execute(w, struct{ Success bool }{true})
	
}