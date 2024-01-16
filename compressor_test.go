package compressor

import (
	"testing"
)

func TestCompressionAndDecompression(t *testing.T) {
	// Data to compress
	data := "Hellooooooooooo, World!"

	// Compress the data
	var compressed []uint8
	compressedSize := Compress([]uint8(data), &compressed)

	// Decompress the data
	decompressed := make([]uint8, len(data))
	decompressedSize := Decompress(compressed, decompressed)

	// Print original, compressed, and decompressed sizes
	t.Logf("Original Size: %d bytes\n", len(data))
	t.Logf("Compressed Size: %d bytes\n", compressedSize)
	t.Logf("Decompressed Size: %d bytes\n", decompressedSize)

	// Print original and decompressed data
	t.Logf("Original Data: %s\n", data)
	t.Logf("Decompressed Data: %s\n", string(decompressed))

	// Check assertions
	if compressedSize >= decompressedSize {
		t.Errorf("Assertion failed: compressedSize should be less than decompressedSize")
	}

	if int(decompressedSize) != len(data) {
		t.Errorf("Assertion failed: decompressedSize should be equal to len(data)")
	}

	if string(decompressed) != data {
		t.Errorf("Assertion failed: string(decompressed) should be equal to data")
	}
}
