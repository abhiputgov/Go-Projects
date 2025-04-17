package conversion

import (
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	returnable := make([]float64, len(strings))
	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, err
		}
		returnable[stringIndex] = floatPrice
	}

	return returnable, nil
}
