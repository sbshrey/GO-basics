package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate User Input
// ==> Show error message & exit if invalid input is provided
// ==> No negative numbers
// ==> Not 0
// 2) Store the result into file

const accountBalanceFile = "results.txt"

func writeTextToFile(optText string) {
	os.WriteFile(accountBalanceFile, []byte(optText), 0644)
}

func main() {
	revenue, errRevenue := inputText("Revenue: ")
	if errRevenue != nil {
		fmt.Println(errRevenue)
		// return
		panic(errRevenue)
	}
	expenses, errExpenses := inputText("Expenses: ")
	if errExpenses != nil {
		fmt.Println(errExpenses)
		return
	}
	taxRate, errTaxRate := inputText("Tax Rate: ")
	if errTaxRate != nil {
		fmt.Println(errTaxRate)
		return
	}

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)
	optText := outputText(ebt, profit, ratio)
	writeTextToFile(optText)
}

func inputText(inputText string) (float64, error) {
	var userInput float64
	fmt.Print(inputText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("input should be positive")
	}

	return userInput, nil
}

func calculateValues(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio // personally preferred to know what i am returning
	// return
}

func outputText(earningsBeforeTax float64, profit float64, ratio float64) string {
	formattedEBT := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", earningsBeforeTax, profit, ratio)
	fmt.Print(formattedEBT)
	// fmt.Println("Profit: ", profit)
	// fmt.Printf("Ratio: %.2f\n", ratio)
	return formattedEBT
}
