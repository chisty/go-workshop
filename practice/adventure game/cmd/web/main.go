package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	game "github.com/chisty/go-adventure-game"
)

func main() {
	port := flag.Int("port", 3000, "The port to run the Application.")
	fileName := flag.String("file", "gopher.json", "The Story Json File.")
	flag.Parse()

	fmt.Println("Story Json FileName is: ", *fileName)

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Error at file open: ", err.Error())
	}

	story, err := game.ParseJsonStory(file)
	if err != nil {
		fmt.Println("Erro in Json Decode: ", err.Error())
	}

	// log.Printf("%+v\n", story)
	fmt.Printf("Starting the server on port %d\n", *port)

	// tpl := template.Must(template.New("").Parse("hello"))
	// h := game.NewHandler(story, game.WithTemplate(tpl))

	h := game.NewHandler(story)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
