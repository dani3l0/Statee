package utils

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Directly reads and trims string from fire
func Cat(target string) (string, error) {
	rawdata, err := os.ReadFile(target)
	if err != nil {
		return "", err
	}
	data := string(rawdata)
	data = strings.Trim(data, "\n")
	data = strings.Trim(data, " ")
	return string(data), nil
}

// Uses the above function and converts result to integer
func CatInt(target string) (int, error) {
	wtf, _ := regexp.Compile("[^0-9]+")
	data, err := Cat(target)
	if err != nil {
		return -127, err
	}

	data = wtf.ReplaceAllString(data, "")
	number, err := strconv.Atoi(data)
	if err != nil {
		return -127, err
	}

	return number, nil
}
