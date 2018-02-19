package io

import "os"

func JointFile(path string, fileFlow chan string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for part := range fileFlow {
		if _, err = f.WriteString(part); err != nil {
			panic(err)
		}
	}
}
