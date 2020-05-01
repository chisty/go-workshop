package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

var (
	key   = []byte("my-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	fmt.Println("Hello")

	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "Request of book with Title %s on Page %s", title, page)
	})

	r.HandleFunc("/foo", middleware(foo))
	r.HandleFunc("/bar", middleware(bar))
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)

	http.ListenAndServe(":3000", r)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-cookie")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-cookie")
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-cookie")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Foo\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Bar\n")
}
