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

	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//THIS IS 100% VOODOO - DONT FUCKING TOUCH THIS UNDER ANY CIRCUMSTANCES
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	

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
