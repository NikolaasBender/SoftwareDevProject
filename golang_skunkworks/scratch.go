package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//ROUTER CONTSRUCTER
//VERY IMPORTANT
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{page}", HomeHandler)
	r.HandleFunc("/", IndexHandle)
	//r.HandleFunc("/second", SecondHandler)
	//r.HandleFunc("/third/{number}", ThirdHandler)
	return r
}

func main() {

	//WE NEED A ROUTER
	r := newRouter()
	//RUNS THE SERVER
	log.Fatal(http.ListenAndServe(":5050", r))
}

type PageVariables struct {
	Date string
	Time string
}

type ContactDetails struct {
	Username string
	Password string
}

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
	t, err := template.ParseFiles(pathVariables["page"])

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, PageVars)

	if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, r.URL.Path[1:])
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("login.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	details := ContactDetails{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// do something with details
	_ = details

	t.Execute(w, struct{ Success bool }{true})
}
