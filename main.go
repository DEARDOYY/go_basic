package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {
	// var accountBalance float64 = 1000.0
	var accountBalance, err = getFloatFromFile(accountBalanceFile, 1000.0)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")

		// panic("Can't continue, sorry.")
	}

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance = accountBalance + depositAmount
			fmt.Println("Balance update! New amount: ", accountBalance)

			writeFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance = accountBalance - withdrawAmount
			fmt.Println("Balance update! New amount: ", accountBalance)

			writeFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank <3")

			return
			// break
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
		// 	accountBalance = accountBalance + depositAmount
		// 	fmt.Println("Balance update! New amount: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdraw amount: ")
		// 	var withdrawAmount float64
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid amount. You can't withdraw more than you have.")
		// 		continue
		// 	}

		// 	accountBalance = accountBalance - withdrawAmount
		// 	fmt.Println("Balance update! New amount: ", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}
	// fmt.Println("Thanks for choosing our bank <3")
}
