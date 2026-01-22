package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
const file="balance.txt"

func main() {
	fmt.Println("Welcome!!!")

	accountBalance,err := read()

	if err!=nil{
		fmt.Println(err)
		panic("System Failed !!!")
	}
	var choice int
	var amount float64

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
			write(accountBalance)
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
			write(accountBalance)
			fmt.Printf("\nYour Money is Withdrawed and your balance is %.2f\n", accountBalance)

		case 4:
			fmt.Println("Thank you. Goodbye.")
			return
		}
	}
}

func write(balance float64){
	text:=fmt.Sprint(balance)
	os.WriteFile(file,[]byte(text),0644)
}

func read()(float64,error){
	data,err:=os.ReadFile(file)
	if(err!=nil){
		return 1000,errors.New("Failed to read File")
	}
	ballanceText:=string(data)
	balance,err:=strconv.ParseFloat(ballanceText,64)
	if(err!=nil){
		return 1000,errors.New("Failed to convert parse data")
	}
	return balance,nil
}

func clearInput() {
	var dump string
	fmt.Scanln(&dump)
} 