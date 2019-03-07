package main

//"fmt"
//"html/template"
//"log"
//"net/http"
//"time"

//"github.com/gorilla/mux"

type PageVariables struct {
	Date string
	Time string
}

type ContactDetails struct {
	Username string
	Password string
}

type Card struct {
	//Add more when we figure out what actually goes in the card
	Text string
}

type BigForm struct{
	nickname string
	email string
	password string
	gender string
	securityQ string
	languages string
	textbox string
}

