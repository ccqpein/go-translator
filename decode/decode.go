package decode

import (
	"bufio"
	"os"
	"strings"
)

// decode ccqline to (key string, symbols []string)
func decodeCCQLine(line string) (string, []string) {
	temp := strings.Split(line, " ")
	key := temp[0][1:]

	temp = temp[1:]
	temp[0] = temp[0][1:]

	lastWord := temp[len(temp)-1]
	lastWord = lastWord[:len(lastWord)-2]
	temp[len(temp)-1] = lastWord

	return key, temp
}

func ReadFile(path string, lineChan chan string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// ccq file sometimes write mutilline of one entry.

	var line []string
	var thisLine string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		thisLine = scanner.Text()
		if thisLine != "" {
			line = append(line, thisLine)
			if thisLine[len(thisLine)-2] == ')' {
				lineChan <- strings.Join(line, " ")
				line = []string{}
			}
		}
	}
}
