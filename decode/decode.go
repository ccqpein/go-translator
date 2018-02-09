package decode

import (
	"bufio"
	"os"
	"strings"
)

type CcqLine struct {
	Symbol  string
	Content []string
}

// decode ccqline to (key string, symbols []string)
func decodeCCQLine(line string) CcqLine {
	temp := strings.Split(line, " ")

	// clean word slice in case some "" inside
	temp = func(a []string) []string {
		b := []string{}
		for _, w := range a {
			if w != "" {
				b = append(b, w)
			}
		}
		return b
	}(temp)

	key := temp[0][1:]

	temp = temp[1:]
	temp[0] = temp[0][1:]

	lastWord := temp[len(temp)-1]
	lastWord = lastWord[:len(lastWord)-2]
	temp[len(temp)-1] = lastWord

	return CcqLine{key, temp}
}

func ReadFile(path string, lineChan chan CcqLine) {
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
				//fmt.Println(line)
				lineChan <- decodeCCQLine(strings.Join(line, " "))
				line = []string{}
			}

			if thisLine[:4] == "#:->" {
				temp := strings.Split(thisLine, " ")
				lineChan <- CcqLine{temp[1], []string{"table change"}}
				line = []string{}
			}

		}
	}

	// close chan
	close(lineChan)
}
