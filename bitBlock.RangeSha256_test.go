package bitBlock

import (
	"bytes"
	"crypto/sha256"
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestRangeSha256(t *testing.T) {
	block := Block{
		buffer: []byte("Hello, World!"),
	}

	// Happy path validation
	t.Run("Happy path", func(t *testing.T) {
		start := 0
		stop := len(block.buffer)

		hash, err := block.RangeSha256(start, stop)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expectedHash := sha256.Sum256(block.buffer[start:stop])
		if !bytes.Equal(hash, expectedHash[:]) {
			t.Fatalf("Expected hash %x, got %x", expectedHash[:], hash)
		}
	})

	// Sad path validations
	t.Run("Bounds check error (stop)", func(t *testing.T) {
		start := 0
		stop := len(block.buffer) + 1

		_, err := block.RangeSha256(start, stop)

		if err == nil || err.Error() != errors.BoundsCheckError {
			t.Fatalf("Expected 'bounds check error', got %v", err)
		}
	})

	t.Run("Bounds check error (start)", func(t *testing.T) {
		start := -1
		stop := len(block.buffer)

		_, err := block.RangeSha256(start, stop)

		if err == nil || err.Error() != errors.BoundsCheckError {
			t.Fatalf("Expected 'bounds check error (start)', got %v", err)
		}
	})

	t.Run("Bounds check error (stop)", func(t *testing.T) {
		start := 0
		stop := -1

		_, err := block.RangeSha256(start, stop)

		if err == nil || err.Error() != errors.BoundsCheckError {
			t.Fatalf("Expected 'bounds check error (stop)', got %v", err)
		}
	})

	t.Run("Stop exceeds start", func(t *testing.T) {
		start := 5
		stop := 4

		_, err := block.RangeSha256(start, stop)

		if err == nil || err.Error() != "stop exceeds start" {
			t.Fatalf("Expected 'stop exceeds start', got %v", err)
		}
	})
}
