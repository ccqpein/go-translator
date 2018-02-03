// using scope table and template to create go source code

package generator

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"../codetemplate"
)

type Function struct {
	FuncName   string
	Parameters string
	ReturnType string
	Body       []string
}

func ReadScope(scopetable map[string][]string, keyword string) []string {
	return scopetable[keyword]
}

func CreateFunc(symbols []string) error {
	thisFunc := Function{}
	f, _ := os.Create("result")
	defer f.Close()

	thisFunc.FuncName = symbols[1]
	// check if this function has return value
	thisFunc.ReturnType = CreateTurpleWithBox(symbols[4])

	thisFunc.Parameters = CreateTurple(symbols[2])
	thisFunc.Body = []string{"aa", "bb"}

	s := codetemplate.GetTemplate("../codetemplate/func.tmpl")
	masterTmpl, err := template.New("function").Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	masterTmpl.Execute(f, thisFunc)

	return nil
}

// Turple used in function parameters, function call, and return values
func CreateTurple(content string) string {
	cutWithSpace := strings.Split(content, " ")
	result := []string{}

	for _, s := range cutWithSpace {
		result = append(result, strings.Replace(s, ":", " ", -1))
	}

	return strings.Join(result, ", ")
}

// Turple with bracket pair
func CreateTurpleWithBox(content string) string {
	return fmt.Sprintf("(%s)", CreateTurple(content))
}
