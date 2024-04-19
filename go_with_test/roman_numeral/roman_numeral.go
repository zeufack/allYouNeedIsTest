package romannumeral

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
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

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}

		// switch {
		// case arabic > 9:
		// 	result.WriteString("X")
		// 	arabic -= 10
		// case arabic > 8:
		// 	result.WriteString("IX")
		// 	arabic -= 9
		// case arabic > 4:
		// 	result.WriteString("V")
		// 	arabic -= 5
		// case arabic > 3:
		// 	result.WriteString("IV")
		// 	arabic -= 4
		// default:
		// 	result.WriteString("I")
		// 	arabic--
		// }
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	// if roman == "III" {
	// 	return 3
	// }
	// if roman == "II" {
	// 	return 2
	// }
	// total := 0
	// for range roman {
	// 	total++
	// }

	var arabic = 0
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return arabic
}
