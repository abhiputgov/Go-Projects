package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceStruct struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (jobToRead TaxIncludedPriceStruct) LoadData() []string {
	file, errOnDataLoad := os.Open("prices.txt")
	if errOnDataLoad != nil {
		fmt.Println("Could not open the file.")
		fmt.Println(errOnDataLoad)
		return []string{}
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Scanning the file failed.")
		fmt.Println(scanner.Err())
		file.Close()
		return []string{}
	}

	return lines

}

func (job TaxIncludedPriceStruct) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f at %.2f", price, job.TaxRate*100)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceStruct {
	return &TaxIncludedPriceStruct{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
