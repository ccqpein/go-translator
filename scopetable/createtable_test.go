package scopetable

import (
	"testing"

	"../decode"
)

var testTable = map[string][]string{}

func TestAddEntry(t *testing.T) {
	AddEntry(&decode.CcqLine{"test", []string{"1", "2"}}, testTable)

	if v1, _ := testTable["test"]; v1[0] != "1" || v1[1] != "2" {
		t.Errorf("add wrong")
	}

	AddEntry(&decode.CcqLine{"test", []string{"a", "b"}}, testTable)

	if v2, _ := testTable["test"]; v2[0] != "a" || v2[1] != "b" {
		t.Errorf("update wrong")
	}

	// make sure there is not two "test" in map
	for k, v := range testTable {
		if k != "test" || v[0] != "a" {
			t.Errorf("data is not consistent")
		}
	}

	t.Log(testTable)
}

func TestCreateTable(t *testing.T) {
	CreateTable("../table.ccq")

	t.Log(ScopeTable)

}
