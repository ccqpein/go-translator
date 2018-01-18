package codetemplate

import "io/ioutil"

// Get template from keyword
func GetTemplate(tempName string) string {
	content, err := ioutil.ReadFile(tempName)
	if err != nil {
		return ""
	}
	return string(content)
}
