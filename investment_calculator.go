package main

import (
	"fmt"
	"math"
)

func main() {
	// camelCase
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64 // must have a type if not given default value

	// fetching data from terminal (scan)
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) // need to pass pointer to work

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println((futureRealValue))
}
