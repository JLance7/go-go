package fileops

import (
	"os"
	"errors"
	"fmt"
	"strconv"
)

func WiteFloatToFile(fileName string, value float64){
	valueText := fmt.Sprint(value) // float to string
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil{
		return 1000, errors.New("Failed to find file")
	}

	valueText := string(data) // convert to string
	value, err := strconv.ParseFloat(valueText, 64) // string to float
	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}
	return value, nil
}