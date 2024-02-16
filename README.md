# Compressor - Simple Repeated Characters Compression Algorithm (Go Version)

This Go package provides a Go implementation of a symmetric compression algorithm designed for sequences containing repeated characters. The algorithm aims to reduce the size of data by efficiently encoding repeated character chunks.

## Installation

To use this package, you can import it into your Go project:

```bash
go get -u github.com/NIR3X/compressor
```

## Usage

Here is an example of how to use the Compressor package in your Go application:

```go
package main

import (
	"fmt"
	"github.com/NIR3X/compressor"
)

func main() {
	// Data to compress
	data := []uint8("Hellooooooooooo, World!")

	// Compress the data
	var compressed []uint8
	compressedSize := compressor.Compress(data, &compressed)

	// Decompress the data
	decompressed := make([]uint8, len(data))
	decompressedSize := compressor.Decompress(compressed, decompressed)

	// Print original, compressed, and decompressed sizes
	fmt.Printf("Original Size: %d bytes\n", len(data))
	fmt.Printf("Compressed Size: %d bytes\n", compressedSize)
	fmt.Printf("Decompressed Size: %d bytes\n", decompressedSize)

	// Print original and decompressed data
	fmt.Printf("Original Data: %s\n", data)
	fmt.Printf("Decompressed Data: %s\n", string(decompressed))
}
```

In this example, the Compressor package compresses the input data, and then decompresses it, demonstrating the basic usage of the compression and decompression functionalities. Adjust the package integration as needed for your specific use case.

## License

[![GNU AGPLv3 Image](https://www.gnu.org/graphics/agplv3-155x51.png)](https://www.gnu.org/licenses/agpl-3.0.html)

This program is Free Software: You can use, study share and improve it at your
will. Specifically you can redistribute and/or modify it under the terms of the
[GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.html) as
published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
