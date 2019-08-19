package romannumerals

import (
	"errors"
	"strings"
)

type romanNumeral struct {
	Value  int
	Symbol string
}

func ToRomanNumeral(i int) (string, error) {
	if i <= 0 || i > 3000 {
		return "", errors.New("number too large")
	}
	romanNumerals := []romanNumeral{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var result strings.Builder
	for _, numeral := range romanNumerals {
		for i >= numeral.Value {
			result.WriteString(numeral.Symbol)
			i -= numeral.Value
		}
	}
	return result.String(), nil
}
