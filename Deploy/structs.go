package main


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
	Posts map[string]Post
}
