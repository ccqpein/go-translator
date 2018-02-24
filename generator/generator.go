// using scope table and template to create go source code

package generator

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"../codetemplate"
	"../scopetable"
)

type Function struct {
	FuncName   string
	Parameters string
	ReturnType string
	Body       []string
}

type Struct struct {
	SName string
	Body  string
}

func GeneratorRouter(f *os.File, startSymbol string, table map[string][]string) (string, error) {
	expression, ok := table[startSymbol]
	if !ok {
		panic("some thing wrong")
	}

	var (
		result string
		err    error
	)

	if keyWTeml, ok := keywords[expression[0]]; ok {
		result, err = keyWTeml(f, expression, scopetable.ScopeTable)
	} else {
		result, err = CreateExpression(expression)
	}

	return result, err
}

func CreateFunc(file *os.File, symbols []string, argv ...interface{}) (string, error) {
	table := argv[0].(map[string][]string)

	thisFunc := Function{}
	length := len(symbols)

	thisFunc.FuncName = symbols[1]
	// check if this function has return value
	if symbols[4][0] == 'G' {
		thisFunc.ReturnType = CreateTurpleWithBox(table[symbols[4]])
	} else {
		thisFunc.ReturnType = symbols[4]
	}

	thisFunc.Parameters = CreateTurple(table[symbols[2]])

	for i := 5; i < length; i++ {
		temp, _ := GeneratorRouter(file, symbols[i], table)
		thisFunc.Body = append(thisFunc.Body, temp)
	}

	s := codetemplate.GetTemplate("func.tmpl")
	masterTmpl, err := template.New("function").Parse(s)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = masterTmpl.Execute(file, thisFunc)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return "", nil
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

func CreateExpression(content []string, argvs ...interface{}) (string, error) {
	var expression string
	var err error = nil

	expression = fmt.Sprintf("%s%s", content[0], CreateTurpleWithBox(content[1:]))

	return expression, err
}

func CreateStruct(file *os.File, a []string, argvs ...interface{}) (string, error) {
	table := argvs[0].(map[string][]string)

	thisStruct := Struct{}

	thisStruct.SName = a[1]
	thisStruct.Body = strings.Replace(CreateTurple(table[a[2]]), ", ", "\n", -1)

	s := codetemplate.GetTemplate("struct.tmpl")
	masterTmpl, err := template.New("struct").Parse(s)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = masterTmpl.Execute(file, thisStruct)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return "", nil

}
