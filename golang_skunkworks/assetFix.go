package main

// import (
// 	"bufio"
// 	//"html/template"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"text/template"
// 	//"reflect"
// 	//"github.com/gorilla/sessions"
// )

// var templateFuncs = template.FuncMap{"rangeStruct": RangeStructer}

// func assetDebug(w http.ResponseWriter, req *http.Request) {
// 	templates := populateTemplates()
// 	requestedFile := req.URL.Path[1:]
// 	template := templates.Lookup()

// 	if template != nil {
// 		template.Execute(w, nil)
// 	} else {
// 		w.WriteHeader(404)
// 	}
// }

// func serveResource(w http.ResponseWriter, req *http.Request) {
// 	path := "public" + req.URL.Path
// 	var contentType string
// 	if strings.HasSuffix(path, ".css") {
// 		contentType = "text/css"
// 	} else if strings.HasSuffix(path, ".png") {
// 		contentType = "image/png"
// 	} else {
// 		contentType = "text/plain"
// 	}

// 	f, err := os.Open(path)

// 	if err == nil {
// 		defer f.Close()
// 		w.Header().Add("Content Type", contentType)

// 		br := bufio.NewReader(f)
// 		br.WriteTo(w)
// 	} else {
// 		w.WriteHeader(404)
// 	}
// }

// func populateTemplates() *template.Template {
// 	result := template.New("templates")

// 	basePath := "templates"
// 	templateFolder, _ := os.Open(basePath)
// 	defer templateFolder.Close()

// 	templatePathRaw, _ := templateFolder.Readdir(-1)

// 	templatePaths := new([]string)
// 	for _, pathInfo := range templatePathRaw {
// 		if !pathInfo.IsDir() {
// 			*templatePaths = append(*templatePaths,
// 				basePath+"/"+pathInfo.Name())
// 		}
// 	}

// 	result.ParseFiles(*templatePaths...)

// 	return result
// }

// // RangeStructer takes the first argument, which must be a struct, and
// // returns the value of each field in a slice. It will return nil
// // if there are no arguments or first argument is not a struct
// func RangeStructer(args ...interface{}) []interface{} {
// 	if len(args) == 0 {
// 		return nil
// 	}

// 	v := reflect.ValueOf(args[0])
// 	if v.Kind() != reflect.Struct {
// 		return nil
// 	}

// 	out := make([]interface{}, v.NumField())
// 	for i := 0; i < v.NumField(); i++ {
// 		out[i] = v.Field(i).Interface()
// 	}

// 	return out
// }
