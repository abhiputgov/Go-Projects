package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err1 := getUserInput("Revenue: \n")
	expenses, err2 := getUserInput("Expenses: \n")
	taxRate, err3 := getUserInput("taxRate: \n")

	if err1 != nil {
		fmt.Print(err1)
		return
	}
	if err2 != nil {
		fmt.Print(err2)
		return
	}
	if err3 != nil {
		fmt.Print(err3)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeOutputToFile(ebt, profit, ratio)

	outputFormattedText("Earnings before tax: %.2f\n", ebt)
	outputFormattedText("Profit after tax: %.2f\n", profit)
	outputFormattedText("Ratio: %.2f\n", ratio)
	outputFormattedText("All three: %.2f %.2f %.2f", ebt, profit, ratio)
}

func fetchInput(input *float64) {
	fmt.Scan(input)
}

func outputText(text string) {
	fmt.Print(text)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	var err error
	outputText(infoText)
	fetchInput(&userInput)
	if userInput <= 0 {
		err = errors.New("No negative or zero numbers please.")
	} else {
		err = nil
	}
	return userInput, err
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func outputFormattedText(text string, value ...interface{}) {
	fmt.Printf(text, value...)
}

func writeOutputToFile(ebt, profit, ratio float64) {
	outputFromCalcs := fmt.Sprint(ebt, profit, ratio)
	os.WriteFile("profitCalcs.txt", []byte(outputFromCalcs), 0644)
}
