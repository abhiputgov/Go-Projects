package prices

import (
	"fmt"
	"go/priceCalculator/conversion"
	"go/priceCalculator/file_manager"
)

type TaxIncludedPriceJob struct {
	IOManager        file_manager.FileManager `json:"-"`
	TaxRate          float64                  `json:"tax_rate"`
	InputPrices      []float64                `json:"input_prices"`
	TaxIncludedPrice map[string]string        `json:"taxed_prices"`
}

func (job *TaxIncludedPriceJob) LoadDataFromFile() {
	lines, errorFromReadFile := job.IOManager.ReadLinesFromFile()
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
	job.IOManager.WriteResult(job)

}

func NewTaxIncludedPriceJob(taxRate float64, fm file_manager.FileManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}
