package main

import (
	"fmt"
)

func main() {
	// camelCase
	const inflationRate = 2.5
	var revenue, expenses, taxRate float64 // must have a type if not given default value

	// fetching data from terminal (scan)
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue) // need to pass pointer to work

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// earningsBeforeTax := revenue - expenses

	// profit := earningsBeforeTax * (1 - taxRate/100)

	// ratio := earningsBeforeTax / profit

	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	// formattedEBT := fmt.Sprintf("EBT: %.2f\n", earningsBeforeTax)
	// fmt.Print(formattedEBT)
	// fmt.Println("Profit: ", profit)
	// fmt.Println("Ratio: ", ratio)
	outputText(ebt, profit, ratio)

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
	fmt.Println("Ratio: ", ratio)
}
