// Read ccq file and create scopetable and dependency table.
package scopetable

import (
	"../decode"
)

//:= TODO: may use sync.Map in futrue
var (
	ScopeTable    = map[string][]string{}
	ScopeDepTable = map[string][]string{}
)

// Switch tables between scopetable and deptable when read line "#:->"
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
// Add entry in table
func AddEntry(line *decode.CcqLine, table map[string][]string) {
	table[line.Symbol] = line.Content
}

// Read ccq file in ccqline struct channel
// Use AddTableRouter to pick which table to add.
func CreateTable(path string) {
	lineChan := make(chan decode.CcqLine)
	var tableFlag *map[string][]string = nil
	var skip bool = false

	go decode.ReadFile(path, lineChan)

	for line := range lineChan {
		//fmt.Println(line)

		if tableFlag, skip = AddTableRouter(&line, tableFlag); !skip {
			AddEntry(&line, *tableFlag)
		}
	}
}
