package propertybased

import (
	"testing"
)

func TestNumConv(t *testing.T) {
	cases := []struct {
		Description string
		Number      int
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

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func BenchmarkArabToRoman(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = ArabToRoman(1776)
	}
}
