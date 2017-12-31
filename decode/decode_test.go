package decode

import (
	"fmt"
	"testing"
	"time"
)

func TestDecodeCCQLine(t *testing.T) {
	line := "(Adam (G546 G549))"

	if key, symbols := decodeCCQLine(line); key != "Adam" ||
		symbols[0] != "G546" ||
		symbols[1] != "G549" {
		fmt.Println(key, symbols)
		t.Errorf("decode wrong")
	}

}

func TestReadFile(t *testing.T) {
	tempchan := make(chan string)

	go ReadFile("./test.ccq", tempchan)
	go func() {
		time.Sleep(time.Second * 2) // close channal after 2 seconds
		close(tempchan)
	}()

	for s := range tempchan {
		t.Log(s)
	}

}
