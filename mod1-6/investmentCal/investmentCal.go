package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6.5
	var investmentAmount, years, expectedReturnRate float64

	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter the number of years: ")
	fmt.Scan(&years)

	outputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years)

	// Example using single value
	formattedFV := formatText("The future value: %.2f\n", futureValue)

	// Example using multiple values
	formattedRFV := formatText("Investment of %.2f will be worth %.2f after inflation\n",
		investmentAmount, futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// Additional example showing multiple parameters
	formattedSummary := formatText("After %.0f years at %.1f%% return rate and %.1f%% inflation:\n", years, expectedReturnRate, inflationRate)
	outputText(formattedSummary)
}

func outputText(text string) {
	fmt.Print(text)
}

func formatText(format string, values ...interface{}) string {
	return fmt.Sprintf(format, values...)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, inflationRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	frv = fv / math.Pow(1+(inflationRate/100), years)

	return fv, frv // (OR) return (OR) nothing here --> this is to be used if the return indicator is not mentioning the return values it is optional
}
