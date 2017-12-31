package decode

import (
	"fmt"
	"testing"
	"time"
)

func TestDecodeCCQLine(t *testing.T) {
	line := "(Adam (G546 G549))"
	line2 := "(G798  (f \"./table.ccq\" :direction :output :if-exists :supersede :if-does-not-exist   :create))"

	if key, symbols := decodeCCQLine(line); key != "Adam" ||
		symbols[0] != "G546" ||
		symbols[1] != "G549" {
		fmt.Println(key, symbols)
		t.Errorf("decode wrong")
	}

	if key, symbols := decodeCCQLine(line2); key != "G798" ||
		symbols[0] != "f" ||
		symbols[1] != "\"./table.ccq\"" {
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
		t.Log(decodeCCQLine(s))
	}

}
