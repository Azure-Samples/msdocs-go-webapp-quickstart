package main

import (
	"html/template"
	"net/http"
)

var name string
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/favicion", favicon)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		name = req.FormValue("name")
		http.Redirect(w, req, "/hello", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "hello.gohtml", name)
}

func favicon(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "/favicon.ico")
}
