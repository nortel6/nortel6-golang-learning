package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(stream []byte) (int, error) {
	r := reader.r
	// Read from the ioReader we wrapped
	// Read size determine by stream passed in
	n, err := r.Read(stream)

	// Modify the stream
	for i := 0; i < n; i++ {
		c := stream[i]
		if c >= 'A' && c <= 'Z' {
			stream[i] = (c-'A'+13)%26 + 'A'
		} else if c >= 'a' && c <= 'z' {
			stream[i] = (c-'a'+13)%26 + 'a'
		}
	}

	// ioReader handles the n and err for us
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
