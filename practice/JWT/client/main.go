package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("my-secret-key")

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := generateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// fmt.Fprintf(w, validToken)

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:9000/", nil)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	req.Header.Add("token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}

func handleRequests() {
	http.HandleFunc("/", handleHomePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "chisty"
	claims["exp"] = time.Now().Add(time.Minute * 30)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("JWT Error: ", err.Error())
	}

	return tokenString, nil
}

func main() {
	fmt.Println("Cliet is running")
	handleRequests()
}
