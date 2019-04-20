package main

//"fmt"
//"html/template"
//"log"
//"net/http"
//"time"

//"github.com/gorilla/mux"

type UserPage struct {
	display_name  string
	display_email string
	display_info  string
	feed          []Post
	tasks         []Task
	pins          []Task
	projects      []Project
}

type ProjectPage struct {
	project_name string
	todo         []Task
	working      []Task
	done         []Task
	users        []User
}

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
type Task struct {
	task_name string
	fill_task string
	Comments  []Comment
}

type Comment struct {
	shmoo string
}

type Feed struct {
	Tasks []Task
}

var postTests = []Task{
	Task{
		Title:   "Test title 1",
		Content: "hi i am content 1",
		Link:    "/post/0",
		Type:    1,
		Comments: []Comment{
			Comment{
				shmoo: "Comment 1 for task 1",
			},
			Comment{
				shmoo: "Comment 2 for task 1",
			},
			Comment{
				shmoo: "Comment 3 for task 1",
			},
			Comment{
				shmoo: "Comment 4 for task 1",
			},
		},
	},
	Task{
		Title:   "Test title 2",
		Content: "hi i am content 2",
		Link:    "/post/1",
		Type:    1,
		Comments: []Comment{
			Comment{
				shmoo: "Comment 1 for task 2",
			},
			Comment{
				shmoo: "Comment 2 for task 2",
			},
			Comment{
				shmoo: "Comment 3 for task 2",
			},
			Comment{
				shmoo: "Comment 4 for task 2",
			},
		},
	},
}
