package main

import(
	"fmt"
)

func main(){
	fmt.Println("Profit Calculator")
	// var revenue, expenses, tax_rate float64
	revenue := outputPerform("Revenue : ")
	expenses := outputPerform("Expenses : ")
	taxRate := outputPerform("tax rate : ")

	calculatePrint(revenue,expenses,taxRate)	
}

func calculatePrint(revenue,expenses,tax_rate float64){
	EBT:=revenue-expenses
	if(EBT<0){
		fmt.Printf("you are in loss of %.2f\n",EBT)
		return
	}
	fmt.Printf("Earning before tax is %.2f \n",EBT)
	Profit:=EBT*(1-tax_rate/100)
	fmt.Printf("Earning After tax is %.2f\n",Profit)
	ratio:=EBT/Profit
	fmt.Printf("Ratio is: %f\n",ratio)
}

func outputPerform (s string)(float64){
	var input float64
	fmt.Print(s)
	// fmt.Scan(&input)
	for {
    _, err := fmt.Scan(&input)
    if err == nil {
        break
    }
    fmt.Println("Please enter a valid number")
	}

	return input
}