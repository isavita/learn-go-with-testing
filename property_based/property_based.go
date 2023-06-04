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
	roman.WriteString(baseRoman(n))
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

func RomanToArabic(s string) int {
	var n int
	for i := 0; i < len(s); i++ {
		currentSymbol := string(s[i])
		currentValue := romanNumeralValue(currentSymbol)
		if i < len(s)-1 {
			nextSymbol := string(s[i+1])
			nextValue := romanNumeralValue(nextSymbol)
			if nextValue > currentValue {
				currentValue = -currentValue
			}
		}
		n += currentValue
	}
	return n
}

func romanNumeralValue(symbol string) int {
	switch symbol {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func findParts(n, part int) int {
	return n / part
}

func RomanToArab(s string) int {
	var n int
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(s, numeral.Symbol) {
			n += numeral.Value
			s = s[len(numeral.Symbol):]
		}
	}
	return n
}

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (numerals RomanNumerals) ValueOf(symbols ...byte) int {
	roman := string(symbols)
	for _, numeral := range numerals {
		if numeral.Symbol == roman {
			return numeral.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
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

func ConvertToRoman(n int) string {
	var roman strings.Builder
	for _, numeral := range allRomanNumerals {
		for n >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}
	return roman.String()
}

func ConvertToArabic(roman string) (n int) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		n += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, numeral := range r {
		if numeral.Symbol == symbol {
			return true
		}
	}
	return false
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		if i+1 < len(w) && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
