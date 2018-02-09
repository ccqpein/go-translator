package main

import (
	"os"

	"./generator"
	"./scopetable"
)

func main() {
	args := os.Args
	//fmt.Println(args)

	scopetable.CreateTable(args[1])
	//fmt.Println(scopetable.ScopeTable)

	testFunc := scopetable.ScopeTable["Adam"][0]

	//fmt.Println(scopetable.ScopeTable[testFunc])
	generator.CreateFunc(scopetable.ScopeTable[testFunc],
		scopetable.ScopeTable)
}
