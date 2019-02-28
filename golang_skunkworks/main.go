package main

import (
	//"assets"
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	//assets "homeMade"

	"github.com/gorilla/mux"

	"github.com/gorilla/sessions"
)

// Templates
var navigationBarHTML string
var homepageTpl *template.Template
var secondViewTpl *template.Template
var thirdViewTpl *template.Template
var indexViewTpl *template.Template
var aboutViewTpl *template.Template


// Config provides basic configuration
type Config struct {
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// HTMLServer represents the web service that serves up HTML
type HTMLServer struct {
	server *http.Server
	wg     sync.WaitGroup
}


func init() {
	//navigationBarHTML = assets.MustAssetString("templates/navigation_bar.html")

	// homepageHTML := assets.MustAssetString("templates/index.html")
	// homepageTpl = template.Must(template.New("homepage_view").Parse(homepageHTML))
	homepageTpl = template.Must(template.ParseFiles("index.html"))

	// secondViewHTML := assets.MustAssetString("templates/second_view.html")
	// secondViewTpl = template.Must(template.New("second_view").Parse(secondViewHTML))

	// // aboutViewHTML := assets.MustAssetString("templates/about.html")
	// // aboutViewTpl = template.Must(template.New("about").Parse(aboutViewHTML))

	// thirdViewFuncMap := ThirdViewFormattingFuncMap()
	// thirdViewHTML := assets.MustAssetString("templates/third_view.html")
	// thirdViewTpl = template.Must(template.New("third_view").Funcs(thirdViewFuncMap).Parse(thirdViewHTML))

}

func main() {
	serverCfg := Config{
		Host:         "localhost:5000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	htmlServer := Start(serverCfg)

	// http.HandleFunc("/secret", secret)
	// http.HandleFunc("/login", login)
	// http.HandleFunc("/logout", logout)

	defer htmlServer.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmt.Println("main : shutting down")
}

// Start launches the HTML Server
func Start(cfg Config) *HTMLServer {
	// Setup Context
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup Handlers
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/second", SecondHandler)
	router.HandleFunc("/third/{number}", ThirdHandler)
	router.HandleFunc("/index", IndexHandler)
	router.HandleFunc("/About", AboutHandler)
	router.HandleFunc("/login", login)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Create the HTML Server
	htmlServer := HTMLServer{
		server: &http.Server{
			Addr:           cfg.Host,
			Handler:        router,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}

	// Add to the WaitGroup for the listener goroutine
	htmlServer.wg.Add(1)

	// Start the listener
	go func() {
		fmt.Printf("\nHTMLServer : Service started : Host=%v\n", cfg.Host)
		htmlServer.server.ListenAndServe()
		htmlServer.wg.Done()
	}()

	return &htmlServer
}

// Stop turns off the HTML Server
func (htmlServer *HTMLServer) Stop() error {
	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Printf("\nHTMLServer : Service stopping\n")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := htmlServer.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := htmlServer.server.Close(); err != nil {
			fmt.Printf("\nHTMLServer : Service stopping : Error=%v\n", err)
			return err
		}
	}

	// Wait for the listener to report that it is closed.
	htmlServer.wg.Wait()
	fmt.Printf("\nHTMLServer : Stopped\n")
	return nil
}

// Render a template, or server error.
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

// Push the given resource to the client.
func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

// Route Handlers

// HomeHandler renders the homepage view template
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

// SecondHandler renders the second view template
func SecondHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
	}
	render(w, r, secondViewTpl, "second_view", fullData)
}

// ThirdHandler renders the third view template
func ThirdHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	push(w, "/static/third_view.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var queryString string
	pathVariables := mux.Vars(r)
	queryNumber, err := strconv.Atoi(pathVariables["number"])
	if err != nil {
		queryString = pathVariables["number"]
	}
	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Number":        queryNumber,
		"StringQuery":   queryString,
	}
	render(w, r, thirdViewTpl, "third_view", fullData)
}

// IndexHandler renders the index view template
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
	}
	render(w, r, indexViewTpl, "index", fullData)
}

// AboutHandler renders the about view template
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
	}
	render(w, r, aboutViewTpl, "about", fullData)
}

//SECURITY
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
