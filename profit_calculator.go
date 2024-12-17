package main

import (
	"fmt"
)

func main() {
	revenue := inputText("Revenue: ")
	expenses := inputText("Expenses: ")
	taxRate := inputText("Tax Rate: ")

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)
	outputText(ebt, profit, ratio)
}

func inputText(inputText string) float64 {
	var userInput float64
	fmt.Print(inputText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateValues(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio // personally preferred to know what i am returning
	// return
}

func outputText(earningsBeforeTax float64, profit float64, ratio float64) {
	formattedEBT := fmt.Sprintf("EBT: %.2f\n", earningsBeforeTax)
	fmt.Print(formattedEBT)
	fmt.Println("Profit: ", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}
