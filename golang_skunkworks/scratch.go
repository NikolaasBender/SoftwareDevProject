package main

import (
	//"html/template"
	"fmt"
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
	// r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	fmt.Println("'" + r.URL.Path[1:] + "'")
	r.HandleFunc("/", HomeHandler)
	//r.PathPrefix("/").Handler(catchAllHandler) YOU NEED HANDLERS FOR OTHER STUFF TO JUSTIFY THIS
	//http.Handle("/", r)
	return r
}

func main() {

	//WE NEED A ROUTER
	r := newRouter()
	fmt.Println("loop?")
	//RUNS THE SERVER
	log.Fatal(http.ListenAndServe(":5050", r))
}

// func general(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, r.URL.Path[1:])
// }

//THE GERNERAL FORM OF A PAGE
type Page struct {
	Title string
	Body  []byte
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println("HOME HANDLER: '" + r.URL.Path[1:] + "'")
	// if strings.Contains(vars["pagedesired"], ".html") == true {
	// 	http.ServeFile(w, r, vars["pagedesired"])
	// }
	http.ServeFile(w, r, r.URL.Path[1:])
}

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, r.URL.Path[1:])
}
