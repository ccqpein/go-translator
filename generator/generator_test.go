package generator

import "testing"

func TestCreateFunc(t *testing.T) {
	CreateFunc([]string{"func", "helloWorld", "a:int b:int", "->", "string int", "G556", "G557"})
}
