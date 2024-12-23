package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName) // missing file will not fail
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}

func main() {
	accountBalance, err := getFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("----------")
		// panic(err)
	}

	fmt.Println("Welcome to GO Bank!")
	for {
		presentOptions()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your Balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your Deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Your Withdrawal: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			} else if withdrawAmount > accountBalance {
				fmt.Println("Invalid Amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeFloatToFile(accountBalanceFile, accountBalance)
		case 4:
			fmt.Println("Thanks for choosing GO Bank!")
			// break // has a different meaning, only brings out of switch
			return
		default:
			fmt.Println("Please enter a valid choice.")
		}
	}

	// 	if choice == 1 {
	// 		fmt.Println("Your Balance is", accountBalance)
	// 	} else if choice == 2 {
	// 		var depositAmount float64
	// 		fmt.Print("Your Deposit: ")
	// 		fmt.Scan(&depositAmount)

	// 		if depositAmount <= 0 {
	// 			fmt.Println("Invalid Amount. Must be greater than 0.")
	// 			continue
	// 		}
	// 		accountBalance += depositAmount
	// 		fmt.Println("Balance updated! New amount:", accountBalance)
	// 	} else if choice == 3 {
	// 		var withdrawAmount float64
	// 		fmt.Print("Your Withdrawal: ")
	// 		fmt.Scan(&withdrawAmount)
	// 		if withdrawAmount <= 0 {
	// 			fmt.Println("Invalid Amount. Must be greater than 0.")
	// 			continue
	// 		} else if withdrawAmount > accountBalance {
	// 			fmt.Println("Invalid Amount. You can't withdraw more than you have.")
	// 			continue
	// 		}

	// 		accountBalance -= withdrawAmount
	// 		fmt.Println("Balance updated! New amount:", accountBalance)
	// 	} else if choice == 4 {
	// 		fmt.Println("Thanks for choosing GO Bank!")
	// 		break
	// 	} else {
	// 		fmt.Println("Please enter a valid choice.")
	// 	}
	fmt.Println(("See you later!"))
}
