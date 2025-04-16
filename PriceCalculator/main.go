package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		var taxIncludedPrice []float64 = make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrice[priceIndex] = (1 + taxRate) * price
		}
		results[taxRate] = taxIncludedPrice
	}

	fmt.Println(results)
}
