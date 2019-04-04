package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// var templates = template.Must(template.ParseGlob("templates/*"))

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	//will eventually make random key generation
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

var debug = true

//THIS HANDLES ANYTHING IN THE TOP DIRECTORY
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if debug == true {
		fmt.Println("Hit HomeHandler")
	}

	pathVariables := mux.Vars(r)
	fmt.Println("HOME HANDLER: '" + pathVariables["page"] + "'" + "'" + r.URL.Path + "'")

	t, _ := template.ParseFiles(pathVariables["page"])

	t.Execute(w, t)
}

//=====================================================================================
//SUPER BASIC INDEX HANDLER
//=====================================================================================
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if debug == true {
		fmt.Println("Hit IndexHandler")
	}

	err := views.Execute(w, "/view/index.html")
	if err != nil {
		log.Fatal("Cannot Get View ", err)
	}
}

//just demo crap
var titles = []string{"t1", "t2", "t3", "t4"}
var contents = []string{"c1", "c2", "c3", "c4"}

//=====================================================================================
//OUR ATTEMPT FOR DEALING WITH CARDS
//=====================================================================================
func CardHandler(w http.ResponseWriter, r *http.Request) {

	if debug == true {
		fmt.Println("Hit CardHandler")
	}

	session, _ := store.Get(r, "login_cookie")
	fmt.Println(session.Values["username"])
	p := Card{Title: "", Content: ""}
	if session.Values["loggedIn"] != true {
		p = Card{Title: "you're not logged in", Content: "Please login"}
	} else {
		p = Card{Title: "you are logged in", Content: "Your stuff is due soon"}
	}
	t, _ := template.ParseFiles("templates/newcard.html")
	t.Execute(w, p)
}

//=====================================================================================
//DEMO FOR DEALING WITH FORMS
//=====================================================================================
func FormHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("formTest.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	details := BigForm{
		nickname:  r.FormValue("nickname"),
		email:     r.FormValue("email"),
		password:  r.FormValue("password"),
		gender:    r.FormValue("gender"),
		securityQ: r.FormValue("securityQuestion"),
		languages: r.FormValue("Languages"),
		textbox:   r.FormValue("textbox"),
	}

	fmt.Println(details)

	t.Execute(w, struct{ Success bool }{true})

}

//=====================================================================================
//VIEW HANDLER
//=====================================================================================
func ViewHandler(w http.ResponseWriter, r *http.Request) {

	if debug == true {
		fmt.Println("Hit ViewHandler")
	}
	session, _ := store.Get(r, "cookie-name")

	if session.Values["authenticated"] != true {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	pathVariables := mux.Vars(r)
	if debug == true {
		fmt.Println("VIEW HANDLER: '" + pathVariables["page"] + "'" + "'" + r.URL.Path + "'")
	}
	
	page := ""

	if strings.Contains(pathVariables["page"], ".html") == true {
		page = pathVariables["page"]
	} else {
		page = pathVariables["page"] + ".html"
	}


	t, _ := template.ParseFiles(page)
	t.Execute(w, nil)
}

//=====================================================================================
//PFFFT I HAVE NO IDEA WHAT THIS IS FOR
//=====================================================================================
func secret(w http.ResponseWriter, r *http.Request) {

	if debug == true {
		fmt.Println("Hit secret")
	}

	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

//=====================================================================================
//THIS IS THE LOGIN HANDLER
//=====================================================================================
func login(w http.ResponseWriter, r *http.Request) {

	if debug == true {
		fmt.Println("Hit login")
		fmt.Println(r.Method)
	}

	t, _ := template.ParseFiles("auth/login.html")
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	session, _ := store.Get(r, "cookie-name")

	//GET LOGIN INFO
	details := ContactDetails{
		Username: r.FormValue("email"),
		Password: r.FormValue("pwd"),
	}

	// Authentication goes here
	// if(password and username) exists in db == true{
	//	session.Values["authenticated"] = true
	//	http.Redirect(w, r, "/view/index.html", http.StatusFound)
	// }else{
	// 	session.Values["authenticated"] = false
	//	t.Execute(w, nil)
	// }

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Values["name"] = r.FormValue("username")
	session.Save(r, w)

	fmt.Println(session.Values["authenticated"])
	fmt.Println(details)

	http.Redirect(w, r, "/view/userpage.html", http.StatusFound)

}

//=====================================================================================
//THE LOG OUT HANDLER
//=====================================================================================
func logout(w http.ResponseWriter, r *http.Request) {

	if debug == true {
		fmt.Println("Hit logout")
	}

	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/view/index.html", http.StatusFound)
}

//=====================================================================================
//THIS DISPLAYS THE CUSTOM 404 PAGE
//=====================================================================================
func notFound(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("404.html")

	t.Execute(w, nil)
}




//=====================================================================================
//NAY SORT OF FEED WILL BE HANDLED WITH THIS
//=====================================================================================
func FeedHandler(w http.ResponseWriter, r *http.Request) {
	//POPULATE THE FEED WITH THE RIGHT POSTS
	if debug == true {
		fmt.Println("Hit FeedHandler")
	}
	session, _ := store.Get(r, "cookie-name")

	if session.Values["authenticated"] != true {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	

}



//=====================================================================================
//ANY SORT OF POST WILL BE HANDLED HERE
//=====================================================================================
func PostHnadler(w http.ResponseWriter, r *http.Request) {
	if debug == true {
		fmt.Println("Hit PostHandler")
	}
	session, _ := store.Get(r, "cookie-name")

	if session.Values["authenticated"] != true {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	pathVariables := mux.Vars(r)

	k := pathVariables["key"]

	//GET THE POST FROM THE DB

	// p := get from db in the post struct postGet(k)

	t, _ := template.ParseFiles("/view/post.html")
	t.Execute(w, p)

}