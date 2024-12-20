package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile) // missing file will not fail
	if err != nil {
		return 1000, errors.New("failed to Read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return balance, nil
}

func main() {
	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("----------")
		return
	}

	fmt.Println("Welcome to GO Bank!")
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
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
