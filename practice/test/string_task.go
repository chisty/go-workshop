package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	fmt.Println(reverse("The quick bròwn 狐 jumped over the lazy 犬"))

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\r\n")
		//result := reg.ReplaceAllString(text, "")
		result := reverse(text)
		fmt.Printf("%s=%s\n", text, result)
	}
}

func reverse(input string) string {
	runes := []rune(input)
	sz := len(runes)

	for i := 0; i < sz/2; i++ {
		runes[i], runes[sz-1-i] = runes[sz-1-i], runes[i]
	}
	return string(runes)
}
