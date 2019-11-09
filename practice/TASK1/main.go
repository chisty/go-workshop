package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println(r.Form)
	r.ParseForm()
	fmt.Println(r.Form)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		// r.ParseForm()
		fmt.Println("UserName: ", r.Form["username"])
		fmt.Println("Password: ", r.Form["password"])

		fmt.Println("UserName: ", r.FormValue("username"))
		fmt.Println("Password: ", r.FormValue("password"))
	}

	// v := url.Values{}
	// v.Set("name", "Ava")
	// v.Add("friend", "Jess")
	// v.Add("friend", "Sarah")
	// v.Add("friend", "Zoe")
	// // v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	// fmt.Println("Encode= ", v.Encode())
	// fmt.Println(v.Get("name"))
	// fmt.Println(v.Get("friend"))
	// fmt.Println(v["friend"])
}
