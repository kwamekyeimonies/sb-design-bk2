package utils

import "strconv"

func ConvertToDouble(inputString string) (float64, error) {

	result, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func ConvertToString(inputFloat float64) string {

	result := strconv.FormatFloat(inputFloat, 'f', -1, 64)
	return result
}
