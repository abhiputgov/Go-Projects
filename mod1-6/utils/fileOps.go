package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultVal float64) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultVal, errors.New("Couldn't read the file. Therefore initialized to default value.")
	}
	valueString := string(data)
	parsedFloatVal, error := strconv.ParseFloat(valueString, 64)
	if error != nil {
		return defaultVal, errors.New("Could'nt convert the string to float.")
	}
	return parsedFloatVal, nil

}

func WriteFloatToFile(fileToWriteTo string, valueToWrite float64) {
	stringValueFromFile := fmt.Sprint(valueToWrite)
	os.WriteFile(fileToWriteTo, []byte(stringValueFromFile), 0644)
}
