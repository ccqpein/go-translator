package scopetable

import (
	"../decode"
)

//:= TODO: may use sync.Map in futrue
var (
	ScopeTable    = map[string][]string{}
	ScopeDepTable = map[string][]string{}
)

func AddTableRouter(line *decode.CcqLine, defaultTable *map[string][]string) *map[string][]string {
	switch line.Symbol {
	case "dependency-table":
		return &ScopeDepTable
	case "scope-table":
		return &ScopeTable
	default:
		return defaultTable
	}
}

//:= MARK: I can use method instead of function
func AddEntry(line *decode.CcqLine, table map[string][]string) {
	table[line.Symbol] = line.Content
}

func CreateTable(path string) {
	lineChan := make(chan decode.CcqLine)
	var tableFlag *map[string][]string = nil
	go decode.ReadFile(path, lineChan)

	for line := range lineChan {
		tableFlag = AddTableRouter(&line, tableFlag)
		AddEntry(&line, *tableFlag)
	}
}
