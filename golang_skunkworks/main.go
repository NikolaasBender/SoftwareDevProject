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
	// r.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("/assets"))))

	//VIEWS SUB ROUTER
	s := r.PathPrefix("/view").Subrouter()
	s.HandleFunc("/", ViewHandler)
	s.HandleFunc("/{page}", ViewHandler)

	//SESSIONS AND STUFF
	r.HandleFunc("/secret", secret)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//CATCH ALL
	//r.HandleFunc("/", IndexHandler)

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
