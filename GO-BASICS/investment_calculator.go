package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value (adjusted for Inflation): %.1f\n", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Printf("Future Value: %.1f\nFuture Real Value (adjusted for Inflation): %.1f", futureValue, futureRealValue)
	// fmt.Println("Future Real Value (adjusted for Inflation)", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}
