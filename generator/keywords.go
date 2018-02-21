package generator

import (
	"os"
	"strings"
)

var keywords = map[string]func(*os.File, []string, ...interface{}) (string, error){}

// avoid initialization loop
func init() {
	keywords["func"] = CreateFunc
	keywords["return"] = func(f *os.File, a []string, argvs ...interface{}) (string, error) {
		result := "return "
		result += strings.Join(a[1:], ", ")

		return result, nil
	}

	keywords["package"] = func(f *os.File, a []string, argvs ...interface{}) (string, error) {
		tempString := strings.Join(a, " ") + "\n\n"
		f.WriteString(tempString)
		return tempString, nil
	}

	keywords["import"] = func(f *os.File, a []string, argvs ...interface{}) (string, error) {
		tempString := a[0] + "(" + strings.Join(a[1:], "\n") + ")" + "\n\n"
		f.WriteString(tempString)
		return "", nil
	}
}
