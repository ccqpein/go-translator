package scopetable

import (
	"../decode"
)

//:= TODO: may use sync.Map in futrue
var (
	ScopeTable    map[string][]string
	ScopeDepTable map[string][]string
)

//:= MARK: I can use method instead of function
func AddEntry(line *decode.CcqLine, table map[string][]string) {
	table[line.Symbol] = line.Content
}

func CreateTable(path string) {
	lineChan := make(chan decode.CcqLine)

	decode.ReadFile(path, lineChan)

	for line := range lineChan {
		AddEntry(&line, ScopeTable)
	}
}
