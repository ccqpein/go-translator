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

	for _, g := range scopetable.ScopeTable["Adam"] {
		//:= TODO: need find a way concatenate different part of file
		generator.GeneratorRouter(g, scopetable.ScopeTable)
	}

	//fmt.Println(scopetable.ScopeTable[testFunc])
	//generator.CreateFunc(scopetable.ScopeTable[testFunc],
	//scopetable.ScopeTable)
}
