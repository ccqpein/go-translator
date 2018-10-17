package generator

import (
	"os"
	"strings"
)

var keywords = map[string]func(*os.File, []string, ...interface{}) (string, error){}

// avoid initialization loop
/*
initialization loop:
	/Users/cchen386/Desktop/go-translator/generator/keywords.go:8:5 keywords refers to
	/Users/cchen386/Desktop/go-translator/generator/generator.go:54:6 CreateFunc refers to
	/Users/cchen386/Desktop/go-translator/generator/generator.go:31:6 GeneratorRouter refers to
	/Users/cchen386/Desktop/go-translator/generator/keywords.go:8:5 keywords
*/
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

	keywords["struct"] = CreateStruct
}
