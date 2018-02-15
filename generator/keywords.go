package generator

import "strings"

var keywords = map[string]func([]string, ...interface{}) (string, error){}

// avoid initialization loop
func init() {
	keywords["func"] = CreateFunc
	keywords["return"] = func(a []string, argvs ...interface{}) (string, error) {
		result := "return "
		result += strings.Join(a[1:], ", ")

		return result, nil
	}
	keywords["package"] = func(a []string, argvs ...interface{}) (string, error) {
		return strings.Join(a, " "), nil
	}
}
