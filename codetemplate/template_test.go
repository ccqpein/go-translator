package codetemplate

import "testing"

func TestGetTemplate(t *testing.T) {
	t.Log(GetTemplate("func.tmpl"))
}
