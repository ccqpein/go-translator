package scopetable

import (
	"fmt"

	"../decode"
)

//:= TODO: may use sync.Map in futrue
var (
	ScopeTable    = map[string][]string{}
	ScopeDepTable = map[string][]string{}
)

func AddTableRouter(line *decode.CcqLine, defaultTable *map[string][]string) (*map[string][]string, bool) {
	switch line.Symbol {
	case "dependency-table":
		return &ScopeDepTable, true
	case "scope-table":
		return &ScopeTable, true
	default:
		return defaultTable, false
	}
}

//:= MARK: I can use method instead of function
func AddEntry(line *decode.CcqLine, table map[string][]string) {
	table[line.Symbol] = line.Content
}

func CreateTable(path string) {
	lineChan := make(chan decode.CcqLine)
	var tableFlag *map[string][]string = nil
	var skip bool = false

	go decode.ReadFile(path, lineChan)

	for line := range lineChan {
		fmt.Println(line)

		if tableFlag, skip = AddTableRouter(&line, tableFlag); !skip {
			AddEntry(&line, *tableFlag)
		}
	}
}
