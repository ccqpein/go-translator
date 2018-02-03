package generator

import "testing"

func TestCreateFunc(t *testing.T) {
	CreateFunc([]string{"func", "helloWorld:string", "G555", "G556", "G557"})
}
