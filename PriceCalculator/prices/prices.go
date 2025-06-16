package prices

import (
	"fmt"
	"go/priceCalculator/conversion"
	"go/priceCalculator/file_manager"
)

type TaxIncludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]string
}

func (job *TaxIncludedPriceJob) LoadDataFromFile() {
	lines, errorFromReadFile := file_manager.ReadLinesFromFile("prices.txt")
	if errorFromReadFile != nil {
		DisplayTheError(errorFromReadFile, "Error in reading files")
	}
	prices, errFromStrConversion := conversion.StringToFloat(lines)
	if errFromStrConversion != nil {
		DisplayTheError(errFromStrConversion, "Conversion from string to float failed.")
		return
	}
	job.InputPrices = prices
}

func DisplayTheError(err error, brief string) {
	fmt.Println(brief)
	fmt.Println(err)
}

func (job TaxIncludedPriceJob) Process() {
	job.LoadDataFromFile()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := (1 + job.TaxRate) * price
		result[fmt.Sprintf("%.2f @ %.2f", price, job.TaxRate)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrice = result
	file_manager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
