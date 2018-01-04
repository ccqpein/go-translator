package decode

import (
	"fmt"
	"testing"
	"time"
)

func TestDecodeCCQLine(t *testing.T) {
	line := "(Adam (G546 G549))"
	line2 := "(G798  (f \"./table.ccq\" :direction :output :if-exists :supersede :if-does-not-exist   :create))"

	if line1 := decodeCCQLine(line); line1.Symbol != "Adam" ||
		line1.Content[0] != "G546" ||
		line1.Content[1] != "G549" {
		fmt.Println(line1)
		t.Errorf("decode wrong")
	}

	if line2 := decodeCCQLine(line2); line2.Symbol != "G798" ||
		line2.Content[0] != "f" ||
		line2.Content[1] != "\"./table.ccq\"" {
		fmt.Println(line2)
		t.Errorf("decode wrong")
	}
}

func TestReadFile(t *testing.T) {
	tempchan := make(chan CcqLine)

	go ReadFile("./table.ccq", tempchan)
	go func() {
		time.Sleep(time.Second * 2) // close channal after 2 seconds
		//close(tempchan)
	}()

	for s := range tempchan {
		t.Log(s)
	}

}
