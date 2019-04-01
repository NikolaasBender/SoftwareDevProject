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
	r.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("/assets"))))
	//r.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("/static"))))
	
	a := r.PathPrefix("/static").Subrouter()
	a.HandleFunc("/{page}", StaticHandler)

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

	//404 HANDLEING WITH CUSTOM PAGE
	r.NotFoundHandler = http.HandlerFunc(notFound)

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
