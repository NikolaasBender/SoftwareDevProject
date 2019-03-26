package main

import (
	//"fmt"
	//"html/template"
	"fmt"
	"log"
	"net/http"

	//"time"

	"github.com/gorilla/mux"
)

//ROUTER CONTSRUCTER
//VERY IMPORTANT

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("/assets"))))
	r.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("/static"))))
	

	//VIEWS SUB ROUTER
	s := r.PathPrefix("/view").Subrouter()
	s.HandleFunc("/", ViewHandler)
	s.HandleFunc("/{page}", ViewHandler)

	//SESSIONS AND STUFF
	r.HandleFunc("/secret", secret)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)

	
	//DEFAULT ROUTE WHEN SOMEONE HITS THE SITE
	r.HandleFunc("/", IndexHandler)

	if debug == true {
		fmt.Println("")
	}

	return r
}

func main() {

	//WE NEED A ROUTER
	r := newRouter()

	port := ":5050"

	fmt.Println("go to ->  http://localhost" + port)
	//RUNS THE SERVER
	log.Fatal(http.ListenAndServe(port, r))
}
