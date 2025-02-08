package prices

import (
	"fmt"
	"price-calc-concurrency/conversion"
	"price-calc-concurrency/iomanager"
)

type TaxIncludedPriceJob struct{
	IOManager iomanager.IOManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_price"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		IOManager: io,
		TaxRate: taxRate,
	}
}

func (job TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	result := make(map[string]string)
	err := job.LoadData()
	
	if err != nil {
		errorChan <- err
		return
	}

	for _, priceVal := range job.InputPrices {
		formattedPrice := priceVal * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", priceVal)] = fmt.Sprintf("%.2f", formattedPrice)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func (job *TaxIncludedPriceJob) LoadData() error{
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}
	
	job.InputPrices = prices
	return nil
}