package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("my-secret-key")

func serverHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is running")
}

func handleRequests() {
	http.Handle("/", isAuthorized(serverHomePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header.Get("token"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error in JWT parse")
				}
				return secretKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func main() {
	fmt.Println("Server is running")

	handleRequests()
}
