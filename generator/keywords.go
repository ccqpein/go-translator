package generator

import "strings"

var keywords = map[string]func([]string) string{
	"func": func(a []string) string { return "" },
	"return": func(a []string) string {
		result := "return "
		result += strings.Join(a[1:], ", ")

		return result
	},
}
