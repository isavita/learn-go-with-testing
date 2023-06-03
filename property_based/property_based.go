package propertybased

import (
	"strings"
)

var bases [9]int = [9]int{1000, 900, 500, 400, 100, 90, 50, 40, 10}

func ArabToRoman(n int) string {
	var roman strings.Builder
	for _, part := range bases {
		if n == 0 {
			break
		}
		count := findParts(n, part)
		if count > 0 {
			n -= count * part
			s := baseRoman(part)
			roman.WriteString(strings.Repeat(s, count))
		}
	}
	roman.WriteString(baseRoman(n % 10))
	return roman.String()
}

func baseRoman(n int) string {
	switch n {
	case 0:
		return ""
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	case 10:
		return "X"
	case 40:
		return "XL"
	case 50:
		return "L"
	case 90:
		return "XC"
	case 100:
		return "C"
	case 400:
		return "CD"
	case 500:
		return "D"
	case 900:
		return "CM"
	case 1000:
		return "M"
	}

	return ""
}

func findParts(n, part int) int {
	return n / part
}
