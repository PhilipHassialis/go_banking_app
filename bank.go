package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------------")
	}
	printHeader()

	for {
		printMenu()
		choice := userInput()

		switch choice {
		case 1:
			logBalance(accountBalance)
		case 2:
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount - must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			logBalance(accountBalance)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount - must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount - cannot withdraw more than balance")
				continue
			}

			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			logBalance(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Wrong selection")
		}
	}
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 0, errors.New("Failed to read balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse balance")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	var balanceText = fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), fs.ModePerm)
}

func userInput() int {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	return choice
}

func logBalance(accountBalance float64) {
	fmt.Printf("Balance: %.2f\n", accountBalance)
}

func printHeader() {
	fmt.Println("Welcome to Go bank")
	fmt.Println("----------------------------------")
}

func printMenu() {

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
