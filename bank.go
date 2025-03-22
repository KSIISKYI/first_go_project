package main

import (
	"fmt"
)

var balance float64 = 1000.0

func main() {
	fmt.Println(makeTextColor("Welcome to the Gay Bank", TextColorPink))

	user, err := showAuthMenu()

	if err != nil {
		fmt.Println(makeTextColor(err.Error(), TextColorRed))
	} else if user != nil {
		showUserMenu(user)
	}
}

func showAuthMenu() (*User, error) {
	for {
		var choice int

		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			return login(), nil
		case 2:
			return register()
		case 3:
			sayBye()
			return nil, nil
		default:
			fmt.Println(makeTextColor("Invalid choice. Try again.", TextColorRed))
		}
	}
}

func showUserMenu(user *User) {
	for {
		switch scanAndReturnUserChoice() {
		case 1:
			showBalance(user)
		case 2:
			depositMoney(user)
		case 3:
			withdrawMoney(user)
		case 4:
			sayBye()
			return
		default:
			fmt.Println(makeTextColor("Invalid choice. Try again.", TextColorRed))
		}
	}
}

func scanAndReturnUserChoice() (choice int) {
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	return choice
}

// Show balance with blue color
func showBalance(user *User) {
	fmt.Println(makeTextColor(fmt.Sprintf("Balance: %.2f", user.Balance), TextColorBlue))
}

func depositMoney(user *User) {
	var amount float64

	for {
		scanAndSet(makeTextColor("Enter amount to deposit: ", TextColorYellow), &amount)

		if amount < 0 {
			fmt.Println(makeTextColor("Amount cannot be negative", TextColorRed))
		} else {
			user.Balance += amount
			err := updateUser(user)

			if err != nil {
				fmt.Println(makeTextColor(err.Error(), TextColorRed))
			} else {
				fmt.Println(makeTextColor(fmt.Sprintf("Successfully deposited %.2f", amount), TextColorGreen))
			}

			return
		}
	}
}

func withdrawMoney(user *User) {
	var withdrawalAmount float64

	for {
		scanAndSet(makeTextColor("Enter amount to withdraw: ", TextColorYellow), &withdrawalAmount)

		if withdrawalAmount < 0 {
			fmt.Println(makeTextColor("Amount cannot be negative", TextColorRed))
		} else if withdrawalAmount > user.Balance {
			fmt.Println(makeTextColor("Insufficient balance", TextColorRed))
		} else {
			user.Balance -= withdrawalAmount
			err := updateUser(user)

			if err != nil {
				fmt.Println(makeTextColor(err.Error(), TextColorRed))
			} else {
				fmt.Println(makeTextColor(fmt.Sprintf("Successfully withdrew %.2f", withdrawalAmount), TextColorGreen))
			}

			return
		}
	}
}

// Say bye with green color
func sayBye() {
	fmt.Println(makeTextColor("Bye", TextColorGreen))
}

func scanAndSet(promt string, variable *float64) {
	fmt.Print(promt)
	fmt.Scanln(variable)
}
