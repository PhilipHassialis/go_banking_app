package main

import "fmt"

func main() {

	accountBalance := 1000.0

	fmt.Println("Welcome to Go bank")
	fmt.Println("----------------------------------")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		logBalance(accountBalance)
	} else if choice == 2 {
		fmt.Print("Deposit amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount - must be greater than 0")
			return
		}

		accountBalance += depositAmount
		logBalance(accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdraw amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount - must be greater than 0")
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Invalid amount - cannot withdraw more than balance")
		}

		accountBalance -= withdrawAmount
		logBalance(accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}

func logBalance(accountBalance float64) {
	fmt.Printf("Balance: %.2f\n", accountBalance)

}
