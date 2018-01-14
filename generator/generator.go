// using scope table and template to create go source code

package generator

func ReadScope(scopetable map[string][]string, keyword string) []string {
	return scopetable[keyword]
}

func ReadPredicate()
