package main

import (
	"html/template"
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
	r.HandleFunc("/", general).Methods("GET")
	//THIS IS THE NEW METHOD THAT WORKS
	//TO FUTURE NICK: SEE IF YOU CAN PUT THIS INTO A FUNCTION
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {

	//WE NEED A ROUTER
	r := newRouter()
	//RUNS THE SERVER
	log.Fatal(http.ListenAndServe(":8080", r))
}

func general(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

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
func rawHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

//tHIS LOADS A .TXT FILE FOR VIEWING IN HTML
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
