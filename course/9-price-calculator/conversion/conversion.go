package conversion

import (
	"strconv"
)

func StringsToFloats(stringValues []string) ([]float64, error) {
	floats := make([]float64, len(stringValues))
	for i, stringVal := range stringValues {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, err
		}
		floats[i] = floatVal
	}
	return floats, nil
}