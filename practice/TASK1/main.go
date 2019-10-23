package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//findMaxLen()
	//distributeCoins()
	testChannel()
}

func testChannel() {
	c := make(chan string)
	go show(c)

	fmt.Print(">")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if strings.HasPrefix(input, "stop") {
			fmt.Println("Closing")
			os.Exit(0)
		}
		c <- input
	}
}

func show(c chan string) {
	for i := range c {
		fmt.Print("Showing value: ", i)
		fmt.Print(">")
	}
}

func distributeCoins() {
	for _, value := range users {
		usedCoin := getCoinForUser(value)
		if usedCoin > 10 {
			usedCoin = 10
		}
		distribution[value] = usedCoin
		coins -= usedCoin
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}

func getCoinForUser(name string) int {
	usedCoin := 0
	for _, c := range name {
		switch c {
		case 'a', 'e', 'A', 'E':
			usedCoin++
		case 'i', 'I':
			usedCoin += 2
		case 'o', 'O':
			usedCoin += 3
		case 'u', 'U':
			usedCoin += 4
		}
	}
	return usedCoin
}

func findMaxLen() {
	maxLen := -1
	for _, value := range names {
		if len := len(value); len > maxLen {
			maxLen = len
		}
	}

	result := make([][]string, maxLen)
	for _, value := range names {
		len := len(value)
		result[len-1] = append(result[len-1], value)
	}

	fmt.Println(result)
}

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))

	names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt", "Emma", "Isabella", "Emily", "Madison",
		"Ava", "Olivia", "Sophia", "Abigail", "Elizabeth", "Chloe", "Samantha", "Addison", "Natalie", "Mia", "Alexis"}
)
