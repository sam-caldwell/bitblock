package bitBlock

import "sync"

// Block - An atomic block unit.
//
//	The Block allows a sequence of related bytes to be
//	read from a file into a single struct for manipulation.
type Block struct {
	lock   sync.Mutex
	buffer []byte
}
