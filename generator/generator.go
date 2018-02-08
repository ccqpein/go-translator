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

func CreateFunc(symbols []string, table map[string][]string) error {
	thisFunc := Function{}
	length := len(symbols)
	f, _ := os.Create("result")
	defer f.Close()

	thisFunc.FuncName = symbols[1]
	// check if this function has return value
	if symbols[4][0] == 'G' {
		thisFunc.ReturnType = CreateTurpleWithBox(table[symbols[4]])
	} else {
		thisFunc.ReturnType = symbols[4]
	}

	thisFunc.Parameters = CreateTurple(table[symbols[2]])

	for i := 5; i < length; i++ {
		thisFunc.Body = append(thisFunc.Body, CreateExpression(table[symbols[i]]))
	}

	s := codetemplate.GetTemplate("../codetemplate/func.tmpl")
	masterTmpl, err := template.New("function").Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	masterTmpl.Execute(f, thisFunc)

	return nil
}

// Turple used in function parameters, function call, and return values
func CreateTurple(content []string) string {
	//cutWithSpace := strings.Split(content, " ")
	result := []string{}

	for _, s := range content {
		result = append(result, strings.Replace(s, ":", " ", -1))
	}

	return strings.Join(result, ", ")
}

// Turple with bracket pair
func CreateTurpleWithBox(content []string) string {
	return fmt.Sprintf("(%s)", CreateTurple(content))
}

func CreateExpression(content []string) string {
	var expression string
	//:= MARK: stop here, need design keyword map
	if keyWTeml, ok := keywords[content[0]]; ok {
		expression = keyWTeml(content)
	} else {
		expression = fmt.Sprintf("%s%s", content[0], CreateTurpleWithBox(content[1:]))
	}

	return expression
}
