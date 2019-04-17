package main

//"fmt"
//"html/template"
//"log"
//"net/http"
//"time"

//"github.com/gorilla/mux"

type ContactDetails struct {
	Username string
	Password string
}

type BigForm struct {
	nickname  string
	email     string
	password  string
	gender    string
	securityQ string
	languages string
	textbox   string
}

//WE WILL WANT TO GET MORE STUFF ABOUT EACH POST
type Post struct {
	Title   string
	Content string
	Link    string
	Type    int
}

type Feed struct {
	Title string
	Posts []Post
}

var postTests = []Post{
	Post{
		Title:   "Test title 1",
		Content: "hi i am content 1",
		Link:    "/post/0",
		Type:    1,
	},
	Post{
		Title:   "Test title 2",
		Content: "hi i am content 2",
		Link:    "/post/1",
		Type:    1,
	},
}
