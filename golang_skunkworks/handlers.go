package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"time"

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
	session.Values["logedIn"] = true
	session.Save(r, w)

	// do something with details
	fmt.Println(session.Values["username"])

	t.Execute(w, struct{ Success bool }{true})
}

//var okCookie =

func CardHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	session, _ := store.Get(r, "cookie_name")
	p := Card{Title: "", Content: ""}
	if session.Values["loggedIn"] != true {
		p = Card{Title: "you're not logged in", Content: "Please login"}
	} else {
		p = Card{Title: "you are logged in", Content: "Your stuff is due soon"}
	}
	t, _ := template.ParseFiles("templates/newcard.html")
	t.Execute(w, p)
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

var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

func templateStackTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	cards := []Card{
		{"t1", "c1"},
		{"t2", "c2"},
		{"t3", "c3"},
		{"t4", "c4"},
		{"t5", "c5"},
		{"t6", "c6"},
	}

	t, err := template.ParseFiles("templates/fillTable.html").Funcs(templateFuncs)
	// /t, err := t.ParseFiles("templates/fillTable.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, cards)
	if err != nil {
		panic(err)
	}
}

// RangeStructer takes the first argument, which must be a struct, and
// returns the value of each field in a slice. It will return nil
// if there are no arguments or first argument is not a struct
func RangeStructer(args ...interface{}) []interface{} {
	if len(args) == 0 {
		return nil
	}

	v := reflect.ValueOf(args[0])
	if v.Kind() != reflect.Struct {
		return nil
	}

	out := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		out[i] = v.Field(i).Interface()
	}

	return out
}
