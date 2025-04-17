package prices

import (
	"bufio"
	"fmt"
	"go/priceCalculator/conversion"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate     float64
	InputPrices []float64
}

func (job *TaxIncludedPriceJob) LoadDataFromFile() {
	file, err := os.Open("prices.txt")
	if err != nil {
		displayTheError(err, "Couldn't open file!")
		return
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errFromScanner := scanner.Err()
	if errFromScanner != nil {
		displayTheError(errFromScanner, "Reading from line ended.")
		file.Close()
		return
	}
	prices, errFromStrConversion := conversion.StringToFloat(lines)
	if errFromStrConversion != nil {
		displayTheError(errFromStrConversion, "Conversion from string to float failed.")
	}
	job.InputPrices = prices
}

func displayTheError(err error, brief string) {
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
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
