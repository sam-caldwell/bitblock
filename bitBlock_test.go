package bitBlock

import "testing"

func TestBlockStruct(t *testing.T) {
	block := Block{}
	if block.buffer != nil {
		t.Fatal("expect block buffer to be nil by default")
	}
}
