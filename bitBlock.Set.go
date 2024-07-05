package bitBlock

// Set - Load data into the buffer
func (block *Block) Set(b []byte) {
	block.lock.Lock()
	defer block.lock.Unlock()

	if b == nil {
		b = make([]byte, 0)
	}

	if block.buffer == nil {
		block.buffer = make([]byte, len(b))
	}

	copy(block.buffer, b)
}
