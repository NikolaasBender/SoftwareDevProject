package main

import (
	//"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//ROUTER CONTSRUCTER
//VERY IMPORTANT
func newRouter() *mux.Router {
	r := mux.NewRouter()
	//THIS IS FROM THE OLD METHOD...THE BUILT IN METHOD
	//r.HandleFunc("/", general).Methods("GET")
	//THIS IS THE NEW METHOD THAT WORKS
	//TO FUTURE NICK: SEE IF YOU CAN PUT THIS INTO A FUNCTION
	// staticFileDirectory := http.Dir("./assets/")
	// staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/", HomeHandler)
	//YOU'RE GONNA WANNA CHANGE THESE. YOU MIGHT WANT PROJECTS TO BE UNDER USERS
	//r.HandleFunc("/projects/{project}", ProjectHandler)
	//r.HandleFunc("/user", UserHandler)

	http.Handle("/", r)
	return r
}

func main() {

	//WE NEED A ROUTER
	r := newRouter()
	//RUNS THE SERVER
	log.Fatal(http.ListenAndServe(":8080", r))
}

// func general(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, r.URL.Path[1:])
// }

//THE GERNERAL FORM OF A PAGE
type Page struct {
	Title string
	Body  []byte
}

//THIS MIGHT NEED TO BE REMOVED
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//THIS JUST DISPLAYS A TEXT FILE ON THE WEBPAGE
//DEFINITELY USEFUL FOR STUFF - MAYBE IT CAN WORK WITH CSV?
// func rawHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	renderTemplate(w, "view", p)
// }

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, r.URL.Path[1:])
}
