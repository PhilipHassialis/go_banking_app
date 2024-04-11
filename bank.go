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

	accountBalance, err := getFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------------")
		panic(err)
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
			writeFloatToValue(accountBalance, accountBalanceFile)
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
			writeFloatToValue(accountBalance, accountBalanceFile)
			logBalance(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Wrong selection")
		}
	}
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("failed to read file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}
	return value, nil
}

func writeFloatToValue(value float64, fileName string) {
	var valueText = fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), fs.ModePerm)
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
