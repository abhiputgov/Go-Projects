package main

import (
	"fmt"
	"go/priceCalculator/file_manager"
	"go/priceCalculator/prices"
)

func main() {
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		fm := file_manager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, *fm)
		go priceJob.Process(doneChans[index], errorChans[index])
	}
	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
