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
		accountBalance += depositAmount
		logBalance(accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdraw amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		logBalance(accountBalance)
	}
}

func logBalance(accountBalance float64) {
	fmt.Printf("Balance: %.2f\n", accountBalance)

}
