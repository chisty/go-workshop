package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readInput()
}

func minimumBribes(q []int32) {
	total := 0
	for i := len(q) - 1; i > -1; i-- {
		if q[i] == int32(i+1) {
			continue
		}

		if i-1 >= 0 && q[i-1] == int32(i+1) {
			total++
			q[i], q[i-1] = q[i-1], q[i]
		} else if i-2 >= 0 && q[i-2] == int32(i+1) {
			total += 2
			q[i-2], q[i-1], q[i] = q[i-1], q[i], int32(i+1)
		} else {
			fmt.Println("Too chaotic")
			return
		}
	}

	fmt.Println(total)

}

func diff(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}

func readInput() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var input []int32
		words := strings.Fields(scanner.Text())
		for _, num := range words {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			input = append(input, int32(val))
		}

		minimumBribes(input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
