package generator

import "testing"

/*
func TestCreateFunc(t *testing.T) {
	CreateFunc(nil, []string{"func", "helloWorld", "G590", "->", "G591", "G592", "G593"}, map[string][]string{
		"G590": []string{"s:string"},
		"G591": []string{"string", "int"},
		"G592": []string{"fmt.Println", "\"hello world\"", "s"},
		"G593": []string{"return", "waahaha", "1"},
	})
}*/

func TestCreateTurple(t *testing.T) {
	t.Log(CreateTurple([]string{"string", "int"}))
	t.Log(CreateTurple([]string{"s:string", "a:int"}))
	t.Log(CreateTurple([]string{"A:string", "B:string", "C:int"}))
}

/*
func TestCreateStruct(t *testing.T) {
	t.Log(CreateStruct(nil, []string{"A:string", "B:string", "C:int"}, nil))

}*/
