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
	fmt.Println("Console Tax Calculator")
	fileName := "tax_bracket.txt"
	brackets, err := readTaxBracket(fileName)
	if err != nil {
		log.Fatal("Cannot read the tax file: ", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input Monthly Salary: ")
	salInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Cannot read the input: ", err)
	}

	fmt.Print("Input Yearly Tax Free Allowance (250000): ")
	allowanceInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Cannot read the allowance input: ", err)
	}

	salInput = strings.TrimRight(salInput, "\r\n")
	allowanceInput = strings.TrimRight(allowanceInput, "\r\n")

	salary, err := strconv.ParseFloat(salInput, 64)
	if err != nil {
		log.Fatal("Invalid salary input. ", err)
	}

	allowance, err := strconv.ParseFloat(allowanceInput, 64)
	if err != nil {
		log.Fatal("Invalid allowance input. ", err)
	}

	taxableIncome := salary*12 - allowance
	tax := calculateTax(taxableIncome, brackets)

	fmt.Printf("Salary Per Month %.2f.\nSalary Per Year %.2f\n", salary, salary*12)
	fmt.Printf("Total Allowance: %.2f\nTotal Taxable Income: %.2f\n", allowance, taxableIncome)

	fmt.Printf("Tax Per Year %.2f\nTax Per Month Tax %.2f\n", tax, tax/12)

	for i := 0; i < len(brackets); i++ {
		if brackets[i].Rate != 0 && brackets[i].Tax == 0 {
			break
		}
		fmt.Printf("%d. Amount: %.2f  Rate: %.2f  Tax: %.2f\n", brackets[i].Index, brackets[i].Amount, brackets[i].Rate, brackets[i].Tax)
	}

	fmt.Println("Process complete")

	reader.ReadBytes('\n')
}

func calculateTax(salary float64, taxBracket []bracket) float64 {
	tax := 0.0
	bracketTax := 0.0
	for i := 0; i < len(taxBracket); i++ {
		if salary == 0 {
			break
		}
		if salary > taxBracket[i].Amount {
			bracketTax = taxBracket[i].Amount * taxBracket[i].Rate
			salary -= taxBracket[i].Amount
		} else {
			bracketTax = salary * taxBracket[i].Rate
			salary = 0
		}

		taxBracket[i].Tax = bracketTax
		tax += bracketTax
	}

	return tax
}

func readTaxBracket(fileName string) ([]bracket, error) {
	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var brackets []bracket
	scanner := bufio.NewScanner(file)
	index := 1
	for scanner.Scan() {
		text := scanner.Text()
		bracket, err := getBracket(text)
		if err != nil {
			log.Println("Error: ", err)
		} else {
			bracket.Index = index
			brackets = append(brackets, bracket)
		}
		index++
	}

	return brackets, nil
}

func getBracket(text string) (bracket, error) {
	if strings.TrimSpace(text) == "" {
		return bracket{}, nil
	}

	tokens := strings.Split(text, " ")
	if len(tokens) != 5 {
		return bracket{}, nil
	}

	low, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return bracket{}, err
	}

	high, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return bracket{}, err
	}

	tax, err := strconv.ParseFloat(tokens[4], 64)
	if err != nil {
		return bracket{}, err
	}

	return bracket{
		Amount: high - low + 1,
		Rate:   tax / 100,
	}, nil
}

type bracket struct {
	Index  int
	Amount float64
	Rate   float64
	Tax    float64
}
