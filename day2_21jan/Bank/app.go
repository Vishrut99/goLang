package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome!!!")

	accountBalance := 30000.00
	var amount float64
	var choice int
	for {
		fmt.Println("What do you want")
		fmt.Println("1.Check Balance")
		fmt.Println("2.Deposit")
		fmt.Println("3.Withdraw")
		fmt.Println("4.Exit")

		fmt.Print("Your choice ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Enter a number.")
			clearInput()
			continue
		}
		if choice == 1 {
			fmt.Println(accountBalance)
		} else if choice == 2 {
			fmt.Print("Deposit Amount: ")
			_, err := fmt.Scan(&amount)
			if err != nil {
				fmt.Println("Please enter a valid number ")
				continue
			}
			if amount <= 0 {
				fmt.Println("Amount cant be negative or zero. Enter again")
				continue
			}
			accountBalance += amount
			fmt.Printf("\nYour Money is deposited and your balance is %.2f\n", accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdraw Amount: ")
			_, err := fmt.Scan(&amount)
			if err != nil {
				fmt.Println("Please enter a valid number ")
				continue
			}
			if amount <= 0 {
				fmt.Println("Amount cant be negative or zero. Try again after Sometime")
				continue
			}
			if amount > accountBalance {
				fmt.Print("you don't have enough amount in your account.Go check again \n")
				continue
			}
			accountBalance -= amount
			fmt.Printf("\nYour Money is Withdrawed and your balance is %.2f\n", accountBalance)
		} else if choice == 4{
			fmt.Println("Thank You")
		} else {
			fmt.Print("Enter Valid Choice\n")
		}
	}
}

func clearInput() {
	var dump string
	fmt.Scanln(&dump)
}
