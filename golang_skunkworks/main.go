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
	r.Host("lab-co.org")

	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//THIS IS 100% VOODOO - DONT FUCKING TOUCH THIS UNDER ANY CIRCUMSTANCES
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	//ALL THE REST OF THESE ARE FINE TO MESS WITH
	//VIEWS SUB ROUTER
	s := r.PathPrefix("/view").Subrouter()
	s.HandleFunc("/", ViewHandler)
	s.HandleFunc("/{page}", ViewHandler)

	//THE POST HANDLER
	// p := r.PathPrefix("/post").Subrouter()
	// p.HandleFunc("/", PostHandler)
	// p.HandleFunc("/{key}", PostHandler)

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

	port := ":80"

	fmt.Println("go to ->  http://localhost" + port)
	//RUNS THE SERVER
	log.Fatal(http.ListenAndServe(port, r))
}
