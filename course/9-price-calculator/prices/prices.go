package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/iomanager"
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

func (job TaxIncludedPriceJob) Process() error {
	result := make(map[string]string)
	err := job.LoadData()
	if err != nil {
		return err
	}

	for _, priceVal := range job.InputPrices {
		formattedPrice := priceVal * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", priceVal)] = fmt.Sprintf("%.2f", formattedPrice)
	}
	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)
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