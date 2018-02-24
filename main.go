package main

import (
	"fmt"
	"os"

	"./generator"
	"./scopetable"
)

func main() {
	args := os.Args
	//fmt.Println(args)

	scopetable.CreateTable(args[1])
	fmt.Println(scopetable.ScopeTable)

	f, _ := os.Create("result.go")
	defer f.Close()

	for _, g := range scopetable.ScopeTable["Adam"] {
		generator.GeneratorRouter(f, g, scopetable.ScopeTable)
	}

}
