package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fileURL := "https://www.iclarified.com/images/news/58010/275840/275840.jpg"

	downloadFile(fileURL, "test.jpg")
}

func downloadFile(fileURL string, fileName string) {
	resp, err := http.Get(fileURL)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println(err)
	}

}
