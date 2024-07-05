package bitBlock

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// ReadBit - Read the block one bit at a time, returning the bit at position p and return error if
// position p is out of bounds.
func (block *Block) ReadBit(p uint) (bit bool, err error) {
	block.lock.Lock()
	defer block.lock.Unlock()

	if p >= uint(8*len(block.buffer)) {
		return false, fmt.Errorf(errors.IndexOutOfRange)
	}
	//Get the byte index
	i := p / 8

	//Get the bit position in byte i
	b := p % 8

	//get this byte
	thisByte := block.buffer[i]

	mask := byte(1 << b)

	//bit value
	v := thisByte & mask

	//log.Printf("pos: %02d  bit: %02d mask: %08b v: %08b", p, b, mask, v)

	//if bit b is not zero, then return true indicating set bit.
	return v != 0, err
}
