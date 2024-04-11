package fileops

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("failed to read file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}
	return value, nil
}

func WriteFloatToValue(value float64, fileName string) {
	var valueText = fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), fs.ModePerm)
}
