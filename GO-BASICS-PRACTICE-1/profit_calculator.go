package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Print(`Your EBT is: `)
	fmt.Println(ebt)
	fmt.Print(`Your profit is: `)
	fmt.Println(profit)
	fmt.Print(`The ratio: `)
	fmt.Println(ratio)
}
