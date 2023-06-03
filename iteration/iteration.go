package iteration

import "strings"

func Repeat(s string, n int) string {
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteString(s)
	}
	return b.String()
}
