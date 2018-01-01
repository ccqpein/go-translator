package scopetable

//:= TODO: may use sync.Map in futrue
var (
	ScopeTable    map[string][]string
	ScopeDepTable map[string][]string
)

//:= MARK: I can use method instead of function
func AddEntry(key string, value []string, table map[string][]string) {
	table[key] = value
}

func CreateTable(path string) {
	lineChan := make(chan string)

	decode.ReadFile(path, lineChan)

	for line := range lineChan {

	}
}
