// using scope table and template to create go source code

package generator

import (
	"os"
	"strings"
	"text/template"

	"../codetemplate"
)

type Function struct {
	FuncName   string
	Parameters []string
	ReturnType string
}

func ReadScope(scopetable map[string][]string, keyword string) []string {
	return scopetable[keyword]
}

func CreateFunc(symbols []string) error {
	thisFunc := Function{}
	f, _ := os.Create("result")
	defer f.Close()

	temp0 := strings.Split(symbols[1], ":")

	thisFunc.FuncName = temp0[0]
	thisFunc.ReturnType = temp0[1]

	s := codetemplate.GetTemplate("../codetemplate/func.tmpl")
	template.Must(template.New("function").Parse(s)).Execute(f, thisFunc)

	return nil
}
