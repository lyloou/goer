// https://golang.org/doc/articles/wiki/
// $ cd gowiki
// $ go run wiki.go
// open with browser: localhost:8080

package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	"log"
	"fmt"
	"path/filepath"
	"strings"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {

	if _, err := os.Stat("./data"); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./data", os.ModePerm)
		}
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/to_edit", toedit)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	http.ListenAndServe(":8088", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `%s`, "test 一下！")
}

func toedit(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
}

func index(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}
	arr := make([]string, 0)
	for _, f := range files {
		var extension = filepath.Ext(f.Name())
		var name = strings.TrimRight(f.Name(), extension)
		arr = append(arr, name)
	}

	err = templates.ExecuteTemplate(w, "index.html", arr)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	//http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html", "tmpl/index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
