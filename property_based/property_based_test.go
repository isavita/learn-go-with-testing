package propertybased

import (
	"fmt"
	"testing"
	"testing/quick"
)

// My personal implementation
func TestArabToRoman(t *testing.T) {
	cases := []struct {
		Description string
		Number      uint16
		Want        string
	}{
		{"converts 0 to ''", 0, ""},
		{"converts 1 to 'I'", 1, "I"},
		{"converts 3 to 'III'", 3, "III"},
		{"converts 8 to 'VIII'", 8, "VIII"},
		{"converts 10 to 'X'", 10, "X"},
		{"convert 39 to 'XXXIX'", 39, "XXXIX"},
		{"convert 246 to 'CCXLVI'", 246, "CCXLVI"},
		{"convert 1776 to 'MDCCLXXVI'", 1776, "MDCCLXXVI"},
		{"convert 1918 to 'MCMXVIII'", 1918, "MCMXVIII"},
		{"convert 2023 to 'MMXXIII'", 2023, "MMXXIII"},
		{"convert 2421 to 'MMCDXXI'", 2421, "MMCDXXI"},
		{"converts 4 to 'IV'", 4, "IV"},
		{"converts 9 to 'IX'", 9, "IX"},
		{"converts 40 to 'XL'", 40, "XL"},
		{"converts 90 to 'XC'", 90, "XC"},
		{"converts 400 to 'CD'", 400, "CD"},
		{"converts 900 to 'CM'", 900, "CM"},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ArabToRoman(test.Number)
			assertString(t, got, test.Want)
		})
	}
}

func TestRomanToArab(t *testing.T) {
	cases := []struct {
		Description string
		Number      string
		Want        uint16
	}{
		{"converts 0 to ''", "", 0},
		{"converts 1 to 'I'", "I", 1},
		{"converts 3 to 'III'", "III", 3},
		{"converts 8 to 'VIII'", "VIII", 8},
		{"converts 10 to 'X'", "X", 10},
		{"convert 39 to 'XXXIX'", "XXXIX", 39},
		{"convert 246 to 'CCXLVI'", "CCXLVI", 246},
		{"convert 1776 to 'MDCCLXXVI'", "MDCCLXXVI", 1776},
		{"convert 1918 to 'MCMXVIII'", "MCMXVIII", 1918},
		{"convert 2023 to 'MMXXIII'", "MMXXIII", 2023},
		{"convert 2421 to 'MMCDXXI'", "MMCDXXI", 2421},
		{"converts 4 to 'IV'", "IV", 4},
		{"converts 9 to 'IX'", "IX", 9},
		{"converts 40 to 'XL'", "XL", 40},
		{"converts 90 to 'XC'", "XC", 90},
		{"converts 400 to 'CD'", "CD", 400},
		{"converts 900 to 'CM'", "CM", 900},
	}
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := RomanToArabic(test.Number)
			if got != test.Want {
				t.Errorf("got %d want %d", got, test.Want)
			}
		})
	}
}

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
	{Arabic: 1999, Roman: "MCMXCIX"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, _ := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}

	t.Run("errors when called with number greater than 3999", func(t *testing.T) {
		_, err := ConvertToRoman(4000)
		if err == nil {
			t.Error("expected an error, but didn't get any")
		}

		if err != ErrTooBigNumber {
			t.Errorf("got %q want %q", err, ErrTooBigNumber)
		}
	})
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertionFn := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman, _ := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return arabic == fromRoman
	}

	if err := quick.Check(assertionFn, nil); err != nil {
		t.Error("failed checks", err)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func BenchmarkArabToRoman(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ArabToRoman(1776)
	}
}

func BenchmarkRomanToArab(t *testing.B) {
	for i := 0; i < t.N; i++ {
		RomanToArab("MDCCLXXVI")
	}
}

func BenchmarkConvertToRoman(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ConvertToRoman(1776)
	}
}

func BenchmarkConvertToArabic(t *testing.B) {
	for i := 0; i < t.N; i++ {
		ConvertToArabic("MDCCLXXVI")
	}
}
