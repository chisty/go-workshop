package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "csv input file name")
	timeLimit := flag.Int("limit", 30, "time limit for quiz")
	flag.Parse()

	fmt.Println("Hello Quiz Game")

	questions, err := readcsv(*csvFileName)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	success := startquiz(questions, *timeLimit)
	fmt.Printf("Your score %d out of %d", success, len(questions))
}

func startquiz(questions []question, duration int) int {
	success := 0
	timer := time.NewTimer(time.Duration(duration) * time.Second)
	reader := bufio.NewReader(os.Stdin)

	for _, item := range questions {
		fmt.Print(item.q, " = ")

		answerChan := make(chan string)

		go func() {
			input, _ := reader.ReadString('\n')
			answerChan <- strings.TrimSpace(input)
		}()

		select {
		case <-timer.C:
			fmt.Println("\nYour time is up.")
			return success
		case answer := <-answerChan:
			if strings.EqualFold(answer, item.ans) {
				success++
			}
		}
	}
	return success
}

func readcsv(fileName string) ([]question, error) {
	var questions []question

	wd, _ := os.Getwd()
	fmt.Println("Working dir: ", wd)

	path := filepath.Join(wd, "data", fileName)
	fmt.Println("File Path: ", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		questions = append(questions, question{
			line[0],
			line[1],
		})
	}

	return questions, nil
}

type question struct {
	q   string
	ans string
}
