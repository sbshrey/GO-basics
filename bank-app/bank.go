package main

import "fmt"

func main() {
	var accountBalance = 1000.0

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

		if choice == 1 {
			fmt.Println("Your Balance is", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Your Deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else if choice == 3 {
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
		} else if choice == 4 {
			fmt.Println("Thanks for choosing GO Bank!")
			break
		} else {
			fmt.Println("Please enter a valid choice.")
		}
	}
	fmt.Println(("See you later!"))
}
