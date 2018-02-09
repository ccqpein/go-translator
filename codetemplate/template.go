package codetemplate

import (
	"io/ioutil"
	"path"
	"runtime"
)

// Get template from keyword
// Using runtime package to get this folder path
func GetTemplate(tempName string) string {
	_, filename, _, _ := runtime.Caller(0)
	path := path.Dir(filename)
	content, err := ioutil.ReadFile(path + "/" + tempName)
	if err != nil {
		return ""
	}
	return string(content)
}
