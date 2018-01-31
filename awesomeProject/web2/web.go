package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func upload(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method", request.Method)
	if request.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(writer, token)

	} else {
		request.ParseMultipartForm(32 << 20)
		file, handler, err := request.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(writer, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("method: ", request.Method)
	if request.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 36))
		fmt.Println(h, )
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(writer, token)
	} else {
		request.ParseForm()
		fmt.Println("username:", request.Form["username"])
		fmt.Println("password:", request.Form["password"])
		fmt.Println("token:", request.Form["token"])

		username := request.Form.Get("username")
		template.HTMLEscape(writer, []byte("username:"+username))
		fmt.Println(template.HTMLEscapeString(username))

		writer.Write([]byte("\n"))
		t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = t.ExecuteTemplate(writer, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		if err != nil {
			fmt.Println(err)
		}

	}
}
func sayHelloName(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path:", request.URL.Path)
	fmt.Println("scheme:", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])

	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}

	fmt.Fprintf(writer, "Hello %s", "Lou")
}
