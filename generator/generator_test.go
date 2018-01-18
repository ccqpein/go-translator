package generator

import "testing"

func TestCreateFunc(t *testing.T) {
	t.Log(CreateFunc([]string{"func", "helloWorld:string", "G555", "G556", "G557"}))
}
