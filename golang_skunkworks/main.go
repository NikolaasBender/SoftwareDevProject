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
	// r.HandleFunc("/", IndexHandle)
	// r.HandleFunc("/login", LoginHandle)
	// r.HandleFunc("/formHandler", FormHandler)
	// r.HandleFunc("/newCard", CardHandle)
	// r.HandleFunc("/storage/{page}", StorageHandler)
	// //r.HandleFunc("/tables", templateStackTest)
	// //FINAL RESORT TO GET SOMETHING
	// r.HandleFunc("/{page}", HomeHandler)
	// //r.HandleFunc("/assetdebug", assetDebug)
	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//r.HandleFunc("/third/{number}", ThirdHandler)

	//VIEWS SUB ROUTER
	s := r.PathPrefix("/view").Subrouter()
	s.HandleFunc("/", ViewHandler)
	s.HandleFunc("/{page}", ViewHandler)

	//ASSETS SUB ROUTER
	// a := r.PathPrefix("/assets").Subrouter()
	// a.HandleFunc("/", AssetsHandler)
	// a.HandleFunc("/{page}", AssetsHandler)

	//SESSIONS AND STUFF
	r.HandleFunc("/secret", secret)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//CATCH ALL
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
