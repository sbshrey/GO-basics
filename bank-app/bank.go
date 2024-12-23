package main

import (
	"fmt"

	"example.com/bank-app/fileOps"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileOps.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("----------")
		// panic(err)
	}

	fmt.Println("Welcome to GO Bank!")
	fmt.Println("Contact number:", randomdata.PhoneNumber())

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
			fileOps.WriteFloatToFile(accountBalanceFile, accountBalance)
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
			fileOps.WriteFloatToFile(accountBalanceFile, accountBalance)
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
