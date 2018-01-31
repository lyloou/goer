package main

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	beego.Get("/", func(ctx *context.Context) {
		ctx.WriteString("abcdefg")
	})
	go func() {
		server2 := http.Server{
			Handler: beego.BeeApp.Handlers,
			Addr:    ":54321",
		}

		if err := server2.ListenAndServe(); err != nil {
			fmt.Println("start server2 fail", err)
		}
	}()
	//http.HandleFunc("/", nil)
	if err := http.ListenAndServe(":12345", nil); err != nil {
		fmt.Println("start server fail", err)
	}
}

func HttpHandle(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(232)
	writer.Write([]byte("ooook"))
	return

	fmt.Println("", request.Method)
	fmt.Println("", request.URL)
	fmt.Println("", request.URL.Path)
	fmt.Println("", request.RemoteAddr)
	fmt.Println("", request.UserAgent())
	fmt.Println("", request.Header.Get("Accept"))
	fmt.Println("", request.Cookies())
	fmt.Println("Form:", request.FormValue("id"))
	fmt.Println("PostForm:", request.PostForm.Get("id"))

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("read body fail:", err)
		writer.WriteHeader(500)
		return
	}

	fmt.Println("Body:", string(body));
}
func sample5() {
	router := httprouter.New()
	router.GET("/", handle)
	if err := http.ListenAndServe(":12345", router); err != nil {
		fmt.Println("start server fail", err)
	}
}

func sample4() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	if err := http.ListenAndServe(":12345", mux); err != nil {
		fmt.Println("start server fail", err)
	}
}

func handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Hello, %v, %v", r.URL.Path, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %v", r.URL.Path)
}

//
type MyHandler struct {
}

func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/abc" {
		w.Write([]byte("abc"))
		return
	}

	if r.URL.Path == "/xyz" {
		w.Write([]byte("xyz"))
		return
	}

	w.Write([]byte("index"))
}
func sample3() {
	if err := http.ListenAndServe(":12345", MyHandler{}); err != nil {
		fmt.Println("start http fail", err)
	}
}

func smaple2() {
	http.Handle("/", MyHandler{})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		fmt.Println("start http fail:", err)
	}
}

func smaple1() {
	if err := http.ListenAndServe(":12345", nil); err != nil {
		fmt.Println("start http fail:", err)
	}
}
