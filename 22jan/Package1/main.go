package main

import (
	"fmt"
	"Vishrut.com/Packages/fileOps"
)

const file = "balance.txt"

func main() {
	fmt.Println("Welcome!!!")

	accountBalance, err := fileOps.Read(file)

	if err != nil {
		fmt.Println(err)
		panic("System Failed !!!")
	}
	var choice int
	var amount float64

	for {
		printDetails()

		fmt.Print("Your choice ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Enter a number.")
			clearInput()
			continue
		}

		switch choice {
		case 1:
			fmt.Println(accountBalance)
		case 2:
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
			fileOps.Write(accountBalance, file)
			fmt.Printf("\nYour Money is deposited and your balance is %.2f\n", accountBalance)
		case 3:
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
				fmt.Println("you don't have enough amount in your account.Go check again ")
				continue
			}
			accountBalance -= amount
			fileOps.Write(accountBalance, file)
			fmt.Printf("\nYour Money is Withdrawed and your balance is %.2f\n", accountBalance)

		case 4:
			fmt.Println("Thank you. Goodbye.")
			return
		}
	}
}

func clearInput() {
	var dump string
	fmt.Scanln(&dump)
}
