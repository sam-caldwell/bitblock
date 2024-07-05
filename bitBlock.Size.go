package bitblock

// Size - Return the buffer size
func (block *Block) Size() uint {
	block.lock.Lock()
	defer block.lock.Unlock()

	if block.buffer == nil {
		return 0
	}

	return uint(len(block.buffer))
}
