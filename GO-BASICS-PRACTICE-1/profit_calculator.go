package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter the revenue: ")
	// fmt.Scan(&revenue)
	// fmt.Print("Enter the expenses: ")
	// fmt.Scan(&expenses)
	// fmt.Print("Enter the tax rate: ")
	// fmt.Scan(&taxRate)

	revenue := userInput("Enter the revenue: ")
	expenses := userInput("Enter the expenses: ")
	taxRate := userInput("Enter the tax rate: ")

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Print(`Your EBT is: `)
	fmt.Printf("%.1f\n", ebt)
	fmt.Print(`Your profit is: `)
	fmt.Printf("%.1f\n", profit)
	fmt.Print(`The ratio: `)
	fmt.Printf("%.3f\n", ratio)
}

func userInput(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
