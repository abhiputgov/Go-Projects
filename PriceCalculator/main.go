package main

import (
	"fmt"
	"go/priceCalculator/file_manager"
	"go/priceCalculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fm := file_manager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, *fm)
		priceJob.Process()
	}
}
