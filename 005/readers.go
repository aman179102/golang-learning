package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

////////////////////////////////////////////////////
// 1) Basic Reader example
////////////////////////////////////////////////////
func basicReaderDemo() {
	fmt.Println("---- Basic Reader Demo ----")
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) // read 8 bytes at a time
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v b[:n]=%q\n", n, err, b, b[:n])
		if err == io.EOF {
			break
		}
	}
}

////////////////////////////////////////////////////
// 2) MyReader Example (infinite 'A' reader)
////////////////////////////////////////////////////
type MyReader struct{}

// Read fills the slice with 'A's, returns length and no error
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func myReaderDemo() {
	fmt.Println("\n---- MyReader Demo ----")
	r := MyReader{}
	b := make([]byte, 5)
	r.Read(b)
	fmt.Println(string(b)) // prints "AAAAA"
}

////////////////////////////////////////////////////
// 3) ROT13 Reader Example
////////////////////////////////////////////////////
type rot13Reader struct {
	r io.Reader
}

// Read reads from underlying reader and applies ROT13 transformation
func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		c := b[i]
		// ROT13 for uppercase
		if c >= 'A' && c <= 'Z' {
			b[i] = 'A' + (c-'A'+13)%26
		}
		// ROT13 for lowercase
		if c >= 'a' && c <= 'z' {
			b[i] = 'a' + (c-'a'+13)%26
		}
	}
	return n, err
}

func rot13ReaderDemo() {
	fmt.Println("\n---- ROT13 Reader Demo ----")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println()
}

////////////////////////////////////////////////////
// MAIN
////////////////////////////////////////////////////
func main() {
	basicReaderDemo()
	myReaderDemo()
	rot13ReaderDemo()
}
