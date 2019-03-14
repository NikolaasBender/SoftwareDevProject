package main

import (
	"fmt"
	"html/template"
	"net/http"

	//"reflect"
	"time"

	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
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
	http.ServeFile(w, r, "views/index.html")
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("views/login.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	r.ParseForm()

	var email = r.FormValue("email")  // Data from the form
	var pwd = r.FormValue("password") // Data from the form

	dbPwd := "test"              // DB simulation
	dbEmail := "me@project.yolo" // DB simulation

	if email == dbEmail && pwd == dbPwd {
		http.Redirect(w, r, "http://localhost:5050/views/userpage.html", 301)
	} else {
		http.Redirect(w, r, "https://gogle.com", 301)
	}
}

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
