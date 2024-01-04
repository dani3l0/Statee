package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Finds multiple values for specified key
// For 'key: value' line returns just value
func Greps(contents string, find string) ([]string, error) {
	var found_values []string
	splits := strings.Split(contents, "\n")

	for _, v := range splits {
		if strings.Contains(v, find) {
			v = strings.Split(v, ":")[1]
			v = strings.Trim(v, " ")
			found_values = append(found_values, v)
		}
	}

	if len(found_values) == 0 {
		return found_values, errors.New("greps: no occurences found")
	}

	return found_values, nil
}

// Same as above, but converts values to integers
func GrepsInt(contents string, find string) ([]int, error) {
	text, err := Greps(contents, find)
	wtf, _ := regexp.Compile("[^0-9]+")
	var ints []int

	if err != nil {
		return ints, err
	}

	for _, v := range text {
		str := wtf.ReplaceAllString(v, "")
		numeric, err := strconv.Atoi(str)
		if err == nil {
			ints = append(ints, numeric)
		}
	}

	if len(ints) == 0 {
		err = errors.New("greps: no ints found")
	}

	return ints, err
}

// Same as greps, but returns only first occurence
func Grep(contents string, find string) (string, error) {
	result, err := Greps(contents, find)
	if err != nil {
		return "", err
	}
	return result[0], nil
}

// Same as GrepsInt, but return only first occurence
func GrepInt(contents string, find string) (int, error) {
	result, err := GrepsInt(contents, find)
	if err != nil {
		return -127, err
	}
	return result[0], nil
}
