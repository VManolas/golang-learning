// Example 22a Readers
// The io package specifies the io.Reader interface, which represents the read end of a stream of data.
// The Go standard library contains [many implementations](https://golang.org/search?q=Read#Global) of these interfaces,
// including files, network connections, compressors, ciphers, and others.
// The io.Reader interface has a Read method:
// func (T) Read(b []byte) (n int, err error)
// Read populates the given byte slice with data and returns the number of bytes populated and an error value.
// It returns an io.EOF error when the stream ends.
// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	// The example code creates a strings.Reader and consumes its output 8 bytes at a time.
// 	r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)
// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }

// Example 22b Exercise Readers
//  Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
// package main

// import (
// 	"golang.org/x/tour/reader"
// )

// type MyReader struct{}

// // TODO: Add a Read([]byte) (int, error) method to MyReader.
// func (r MyReader) Read(bytes []byte) (int, error) {
// 	// for i := range bytes {
// 	// 	bytes[i] = 65
// 	// }
// 	// return len(bytes), nil
// 	bytes[0] = 'A'
// 	return 1, nil
// }

// func main() {
// 	reader.Validate(MyReader{})
// 	reader.Validate()
// }

// Example 22c Exercise rot13Reader
// A common pattern is an [io.Reader](https://golang.org/pkg/io/#Reader) that wraps another io.Reader, modifying the stream in some way.
// For example, the [gzip.NewReader](https://golang.org/pkg/compress/gzip/#NewReader) function takes an io.Reader
// (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

// Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
// modifying the stream by applying the [rot13](https://en.wikipedia.org/wiki/ROT13) substitution cipher to all alphabetical characters.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		// fmt.Printf("b: %v %q \n", b, b)
		// fmt.Printf("(b - 'A' + 13): %v %q \n", (b - 'A' + 13), (b - 'A' + 13))
		// fmt.Printf("(b-'A'+13)mod26: %v %q \n", (b-'A'+13)%26, (b-'A'+13)%26)
		// fmt.Printf("'A'+((b-'A'+13)mod26): %v %q \n", 'A'+((b-'A'+13)%26), 'A'+((b-'A'+13)%26))
		b = 'A' + ((b - 'A' + 13) % ('Z' - 'A' + 1)) // %26
	} else if b >= 'a' && b <= 'z' {
		// fmt.Printf("b: %v %q \n", b, b)
		// fmt.Printf("(b - 'a' + 13): %v %q \n", (b - 'a' + 13), (b - 'a' + 13))
		// fmt.Printf("(b-'a'+13)mod26: %v %q \n", (b-'a'+13)%26, (b-'a'+13)%26)
		// fmt.Printf("'a'+((b-'a'+13)mod26): %v %q \n", 'a'+((b-'a'+13)%26), 'a'+((b-'a'+13)%26))
		b = 'a' + ((b - 'a' + 13) % ('Z' - 'A' + 1)) // %26
	}
	return b
}

func rot13beta(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		if b >= 'M' {
			b -= 13
		} else {
			b += 13
		}
	} else if b >= 'a' && b <= 'z' {
		if b >= 'm' {
			b -= 13
		} else {
			b += 13
		}
	}
	return b
}

// The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
func (rot *rot13Reader) Read(bytes []byte) (n int, err error) {
	n, err = rot.r.Read(bytes)
	if err == nil {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				bytes[i] = rot13(bytes[i])
			} else {
				bytes[i] = rot13beta(bytes[i])
			}
		}
	}
	return n, err
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
